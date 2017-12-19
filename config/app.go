/**
 * 基本的app配置 包含中间件 配置等
 * File: app.go 
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 6:27:12 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 7:02:17 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

 package Config

 import(
	Gin "github.com/gin-gonic/gin"
	"github.com/xoxo/crm-x/app/Http/Middleware"
 )

 ////////////////Middleware begin//////////////////////
 func SetupGlobalMiddleware(engine *Gin.Engine)  {
	// Global middleware
	engine.Use(middleware.FormatResponse())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	engine.Use(Gin.Recovery())
 }

 func SetupApiMiddleware(router *Gin.RouterGroup)  {
	// router.Use(middleware.FormatResponse())
	// router.Use(Gin.Recovery())
 }

 func SetupAdminMiddleware(router *Gin.RouterGroup)  {
	// router.Use(middleware.FormatResponse())
	// router.Use(Gin.Recovery())
	router.Use(func(context *Gin.Context) {
		context.Next()
	})
 }
 ////////////////Middleware end//////////////////////
