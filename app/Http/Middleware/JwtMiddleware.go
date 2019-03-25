/**
 * File: FormateResponse.go
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 3:35:09 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 5:16:33 pm
 * -----
 * Copyright 2017 - 2027 GOCRAZY, GOCRAZY
 */

package Middleware

import (
	"go-crazy/app/Services/Auth"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func JwtMiddlewareServe() context.Handler {
	return Auth.GetJwtMiddleware().Serve
}

func JwtMiddleware() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()
		// before request
		user := Auth.GetUserByToken(ctx)
		// 默认用户处理 user。id = 0
		if user.ID == 0 {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.Writef("默认用户没有权限")
			ctx.StopExecution()
			return
		}
		ctx.Values().Set("user", user)

		ctx.Next()
		// after request
	}
}
