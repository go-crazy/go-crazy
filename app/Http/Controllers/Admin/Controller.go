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

package AdminController

import (
	"go-crazy/util"
	"os"

	"github.com/kataras/iris"
)

func ShutDown(ctx iris.Context) {
	pid := os.Getpid()
	ps, _ := os.FindProcess(pid)
	ps.Signal(os.Interrupt)
	util.Api_response(ctx, "shutdown")
}
