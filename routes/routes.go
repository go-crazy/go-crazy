/**
 * File: routes.go 路由索引文件
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 5:56:29 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 5:56:35 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

 package Route

 import(
	"github.com/kataras/iris"
	"github.com/go-crazy/go-crazy/Config"
	// "github.com/go-crazy/elastic"
	// jwt "github.com/go-crazy/authentication"
 )

 func SetupRouter(app *iris.Application) *iris.Application {
	// setup global Middleware
	Config.SetupGlobalMiddleware(app)

	// create group api and  admin
	apiGroup := app.Party("/api")
	adminGroup := app.Party("/admin")
	wsGroup := app.Party("/ws")
	
	// wsElastic := app.Party("/elastic")
	// authGroup := app.Group("auth")

	
	// setup router
	SetupWebRouter(app)
	SetupApiRouter(apiGroup)
	SetupAdminRouter(adminGroup)
	// websocket
	SetupWebsocketRouter(wsGroup,app)

	//Elastic
	// Elastic.InitElastic(wsElastic)

	// jwt 
	// jwt.Init(authGroup)

	// setup up Middleware
	Config.SetupApiMiddleware(apiGroup)
	Config.SetupAdminMiddleware(adminGroup)

	// // Serving static files
	app.StaticWeb("/static", "./static")

	// // error page 
	// // 404
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("你是来玩的吗？")
		ctx.HTML("404 你走错了！")
	})

	return app
 }
