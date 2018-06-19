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
	"github.com/go-crazy/elastic"
	// jwt "github.com/go-crazy/authentication"
 )

 func SetupRouter(engine *iris.Application) *iris.Application {
	// setup global Middleware
	Config.SetupGlobalMiddleware(engine)

	// create group api and  admin
	apiGroup := engine.Party("/api")
	adminGroup := engine.Party("/admin")
	wsGroup := engine.Party("/ws")
	
	wsElastic := engine.Party("/elastic")

	// authGroup := engine.Group("auth")

	
	// setup router
	SetupWebRouter(engine)
	SetupApiRouter(apiGroup)
	SetupAdminRouter(adminGroup)
	// websocket
	SetupWebsocketRouter(wsGroup)

	//Elastic
	Elastic.InitElastic(wsElastic)

	// jwt 
	// jwt.Init(authGroup)

	// setup up Middleware
	Config.SetupApiMiddleware(apiGroup)
	Config.SetupAdminMiddleware(adminGroup)

	// // Serving static files
	// engine.Static("/assets", "./static")
	// // router.StaticFS("/more_static", http.Dir("my_file_system"))

	// // error page 
	// // 404
	// engine.NoRoute(func(c *Gin.Context) {
	// 	c.JSON(404, Gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	return engine
 }
