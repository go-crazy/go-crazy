/**
 * File: CommonController.go 通用控制方法
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 6:18:50 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 6:19:24 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */


 package Controller

 import(
	"github.com/kataras/iris"
	// . "github.com/go-crazy/go-crazy/util"
 )

 func Ping(ctx iris.Context)  {
	// Ping test
	ctx.WriteString("pong")
	// Api_response(c,Gin.H{"user": "", "status": "no value"})
 }