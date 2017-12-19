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

import(
	Gin "github.com/gin-gonic/gin"
)

func Api_response(c *Gin.Context,value Gin.H)  {
	c.Set("api_response",value)
}