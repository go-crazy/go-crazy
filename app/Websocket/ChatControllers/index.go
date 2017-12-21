/**
 * 
 * File: index.go
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 3, 2017-12-20 7:47:45 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 3, 2017-12-20 8:27:00 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */


package WsChat

import (
	Gin "github.com/gin-gonic/gin"
)

var hub *Hub
func InitChatHub() {
	hub = newHub()
	go hub.run()
}

func InitChat(c *Gin.Context) {
	serveWs(hub, c.Writer, c.Request)
}