/**
 * File: FormateResponse.go
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 3:35:09 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 5:16:33 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */


 package middleware

 import (
	 Gin "github.com/gin-gonic/gin"
	 "time"
	 "log"
)

 func Logger2() Gin.HandlerFunc {
	return func(c *Gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		
		status := c.Writer.Status()
		log.Println(status,c.GetRawData)

		api_response, exists := c.Get("api_response")
		if(exists){
			meta := Gin.H{"timestamp": time.Now().UnixNano(), "response_time": latency}

			c.JSON(c.Writer.Status(), Gin.H{"data": api_response, "meta": meta})
		}
	}
}
