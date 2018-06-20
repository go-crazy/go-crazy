package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"os/signal"
	"github.com/kataras/iris"
	"github.com/jinzhu/configor"
	"github.com/go-crazy/go-crazy/routes"
	. "github.com/go-crazy/go-crazy/Config"
	"github.com/go-crazy/go-crazy/util/logger"
)

func main() {
	// load config from file
	configor.Load(&Config, ".env.yml")
	
	os.Setenv("GO_ENV",Config.Env)

	// init path
	InitPath()

	// init logger
	InitLogger()

	// init database
	// InitDB()

	// init api 
	app := iris.New()
	// init routers
	Route.SetupRouter(app)
	// pprint 

	printRouter(app)

	//startNormal(app)
	startGracefulShutdown(app)
}

func  printRouter(app *iris.Application) {
	r := app.GetRoutes()
	for _, r := range r {
		logger.Instance().Info("router info ------------ " + r.String())
	}
}

func startNormal(app *iris.Application)  {
	// Listen and Server in Config.Port
	app.Run(iris.Addr( ":"+Config.Port))
}

func startGracefulShutdown(app *iris.Application)  {
	// graceful-shutdown
	var pid = os.Getpid()
	ps, _ := os.FindProcess(pid)

	// shutdown this app
	// todo add Permissions、clear e.t.c
	app.Get("/down", func(ctx iris.Context) {
		ctx.WriteString("OK")
		ps.Signal(os.Interrupt)
    })

	go func() {
		app.Run(iris.Addr( ":"+Config.Port))
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server ...")

	// Recycle
	clearAll()
	logger.Info("------------------    Recycle   ------------------")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		logger.Info(fmt.Sprintf("Server Shutdown: %s\n", err))
	}
	logger.Info("------------------Server exiting------------------")
}

func clearAll()  {
	CloseDB()
}