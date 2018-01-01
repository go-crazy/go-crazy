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
	Gin "github.com/gin-gonic/gin"
	"github.com/go-crazy/go-crazy/app/Http/Controllers"
 )

 func SetupApiRouter(router *Gin.RouterGroup) {
	// Ping test
	router.GET("/ping", Controller.Ping)
 }