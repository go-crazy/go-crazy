/**
 * File: api.go  api 接口路由
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 6:07:41 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 6:07:45 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

 package Route

 import(
	"github.com/kataras/iris"
 )

 func SetupWebRouter(router *iris.Application)  {
	// Ping test
	// router.LoadHTMLGlob("static/templates/*")

	// router.GET("/", func(c *Gin.Context)  {
	// 	c.HTML(200, "index.tmpl", Gin.H{
	// 		"title": "Go-Crazy",
	// 		"msg": "Welcome to Go-Crazy!",
	// 	})
	//  })

	// router.GET("/uuid", func(c *Gin.Context)  {
	// c.HTML(200, "browser-uuid.html", Gin.H{})
	// })
	 
	// router.GET("/w", func(c *Gin.Context)  {
	// 	c.HTML(200, "websocket.html", Gin.H{})
	//  })
 }