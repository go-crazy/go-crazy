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

package Route

import (
	"go-crazy/app/Http/Middleware"

	"github.com/kataras/iris"
)

var (
	App *iris.Application = nil
)

////////////////Middleware begin//////////////////////
func SetupGlobalMiddleware(engine *iris.Application) {
	// Global middleware
	engine.Use(Middleware.FormatResponse())
	engine.Use(Middleware.CorsMiddleware())

	// engine.Use(Middleware.JwtMiddleware())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
}

func SetupApiMiddleware(router iris.Party) {
}

func SetupAdminMiddleware(router iris.Party) {
}

func SetupJwtAuthMiddleware(router iris.Party) {
	router.Use(Middleware.JwtMiddlewareServe())
	router.Use(Middleware.JwtMiddleware())
}

////////////////Middleware end//////////////////////
