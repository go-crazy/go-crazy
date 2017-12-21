/**
 * websocket路由配置
 * File: websocket.go
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 3, 2017-12-20 7:43:36 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 3, 2017-12-20 7:43:39 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

 package Route

 import (
	Gin "github.com/gin-gonic/gin"
	"github.com/xoxo/crm-x/app/Websocket/ChatControllers"
 )


// use websocket in gin  https://github.com/gin-gonic/gin/issues/51
// r := gin.New()
// r.GET("/ws", func(c *gin.Context) {
// 	handler := websocket.Handler(EchoServer)
// 	handler.ServeHTTP(c.Writer, c.Req)
// })
// r.Run(":8080")

 func SetupWebsocketRouter(router *Gin.RouterGroup)  {
	  WsChat.InitChatHub()
	  router.GET("/chat", WsChat.InitChat)
 }