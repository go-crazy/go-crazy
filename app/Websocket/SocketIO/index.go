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


package SocketIO

import (
	"github.com/kataras/iris"

	"github.com/googollee/go-socket.io"
	"github.com/go-crazy/go-crazy/util/logger"

)



func InitSocketIO( app *iris.Application) {

	server, err := socketio.NewServer(nil)
	if err != nil {
		// app.Logger().Fatal(err)
		logger.AppendError("InitSocketIO ",err)
	}

	server.On("connection", func(so socketio.Socket) {
		logger.Instance().Info("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			logger.AppendInfo("emit: %v", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			logger.Instance().Info("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		logger.AppendError("InitSocketIO error ",err)
	})
	// serve the socket.io endpoint.
	app.Any("/socket.io/{p:path}", iris.FromStd(server))
}