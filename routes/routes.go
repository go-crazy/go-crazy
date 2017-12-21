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
	Gin "github.com/gin-gonic/gin"
	"github.com/xoxo/crm-x/Config"
 )

 func SetupRouter(engine *Gin.Engine) *Gin.Engine {
	
	// setup global Middleware
	Config.SetupGlobalMiddleware(engine)

	// create group api and  admin
	apiGroup := engine.Group("api")
	adminGroup := engine.Group("admin")
	wsGroup := engine.Group("ws")

	// setup router
	SetupWebRouter(engine)
	SetupApiRouter(apiGroup)
	SetupAdminRouter(adminGroup)
	// websocket
	SetupWebsocketRouter(wsGroup)

	// setup up Middleware
	Config.SetupApiMiddleware(apiGroup)
	Config.SetupAdminMiddleware(adminGroup)

	// Serving static files
	engine.Static("/assets", "./static")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))

	return engine
 }
