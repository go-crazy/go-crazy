/**
 * File: utils.go 用于封装全局基础方法
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 5:09:02 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 5:09:06 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

package util

import (
	Gin "github.com/gin-gonic/gin"
	"github.com/kataras/iris/context"
)

func Api_response(ctx context.Context, value interface{}) {
	ctx.Values().Set("api_response", value)
}

func Auth(c *Gin.Context, value Gin.H) {
	c.Set("auth_user", value)
}

func GetUser(c *Gin.Context) {
	_, exists := c.Get("auth_user")
	if exists {
		return
	}
}
