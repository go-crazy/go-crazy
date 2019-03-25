/**
 * File: CorsMiddleware.go
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
	"github.com/iris-contrib/middleware/cors"

	"time"

	"github.com/kataras/iris/context"
)

/*
	app := iris.New()
	app.Use(cors.New(opts))
	app.AllowMethods(iris.MethodOptions)
*/
func CorsMiddleware() context.Handler {
	origin := "*"
	opts := cors.Options{
		AllowedOrigins: []string{origin},
		AllowedHeaders: []string{"Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD"},
		ExposedHeaders: []string{"X-Header"},
		MaxAge:         int((24 * time.Hour).Seconds()),
		// Debug:          true,
	}

	return cors.New(opts)
}
