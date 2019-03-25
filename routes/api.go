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

import (
	"go-crazy/app/Http/Controllers"

	"github.com/kataras/iris"
)

func SetupApiRouter(router iris.Party) {
	// Ping test
	router.Get("/ping", Controller.Ping)

	router.Post("/login", Controller.JwtLogin)

	// 增加认证信息
	securedParty := router.Party("")
	SetupJwtAuthMiddleware(securedParty)

	// 获取当前用户信息
	securedParty.Post("/secured", Controller.UserInfo)

}
