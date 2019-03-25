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
	"go-crazy/app/Websocket"

	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	// "github.com/go-crazy/go-crazy/app/Websocket/SocketIO"
)

func SetupWebsocketRouter(router iris.Party, app *iris.Application) {
	//   SocketIO.InitSocketIO(app)

	// iris ws 方案
	router.Get("/test", Websocket.GetInstance().Handler())
	router.Get("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}
