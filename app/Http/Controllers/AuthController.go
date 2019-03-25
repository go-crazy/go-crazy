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

import (
	// "go-crazy/app/Models"
	"go-crazy/app/Services/Auth"
	"go-crazy/app/Services/User"
	"go-crazy/util"

	"github.com/kataras/iris"
)

type LoginInfo struct {
	Name     string
	Password string
}

func JwtLogin(ctx iris.Context) {
	var c LoginInfo
	if err := ctx.ReadJSON(&c); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	user := UserService.GetUserByName(c.Name)
	token := Auth.GetTokenByUser(user)

	data := make(map[string]interface{})
	data["token"] = token
	data["user"] = user

	util.Api_response(ctx, data)
}

func UserInfo(ctx iris.Context) {

	user := ctx.Values().Get("user")

	// user := Auth.GetUserByToken(ctx)

	// if err := ctx.ReadJSON(&c); err != nil {
	// 	ctx.StatusCode(iris.StatusBadRequest)
	// 	ctx.WriteString(err.Error())
	// 	return
	// }

	// user := Model.GetUserByName(c.Name)
	// token := Auth.GetTokenByUser(user)

	// data := make(map[string]interface{})
	// data["token"] = token
	// data["user"] = user

	util.Api_response(ctx, user)
}
