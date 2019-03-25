package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"go-crazy/app"
	"go-crazy/routes"
	"go-crazy/util/logger"

	"github.com/jinzhu/configor"
	"github.com/kataras/iris"
)

func main() {
	// load config from file
	configor.Load(App.Config(), ".env.yml")

	os.Setenv("GO_ENV", App.Config().Env)
	os.Setenv("DEBUG", App.Config().Debug)

	// init path
	App.InitPath()

	// init logger
	App.InitLogger()

	// init database
	App.InitDB()

	logger.Info("DEBUG:" + os.Getenv("DEBUG"))

	defer clearAll()

	// init api
	server := iris.New()

	App.InitIrisApp(server)
	// init routers
	Route.SetupRouter(server)
	// print
	printRouter(server)

	//startNormal(server)
	startGracefulShutdown(server)
}

func printRouter(app *iris.Application) {
	r := app.GetRoutes()
	for _, r := range r {
		logger.Instance().Info("router info : " + r.String())
	}
}

func startNormal(app *iris.Application) {
	// Listen and Server in Config.Port
	app.Run(iris.Addr(":" + App.Config().Port))
}

func startGracefulShutdown(app *iris.Application) {
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
		app.Run(iris.Addr(":" + App.Config().Port))
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

func clearAll() {
	App.ReleaseResources()
}
