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
	"time"

	"github.com/kataras/iris/context"
)

// 使用方法在util中,具体可以参见
// func Api_response(ctx context.Context,value interface{})  {
// 	ctx.Values().Set("api_response",value)
// }

func FormatResponse() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()
		start := time.Now()
		// before request
		ctx.Next()
		// after request
		// time elapsed
		latency := float64(time.Since(start).Nanoseconds()) / 1000000
		// format response
		msg := ctx.Values().Get("api_response")
		if msg != nil {
			meta := make(map[string]interface{})
			meta["timestamp"] = time.Now().Unix()
			meta["response_time"] = latency

			data := make(map[string]interface{})
			data["meta"] = meta
			data["data"] = msg
			ctx.JSON(data)
		}
		if ctx.GetStatusCode() == 500 {
		}
	}
}
