package Elastic

import (
	Gin "github.com/gin-gonic/gin"
	"github.com/xoxo/crm-x/app/Services/Elastic/Routers"
)

func InitElastic(router *Gin.RouterGroup) {

		router.POST("/set", Routers.SetHandler)

		// router.POST("/get", func(c *Gin.Context)  {
		// 	Routers.GetHandler(c.Writer, c.Request)
		// })
		router.POST("/get", Routers.GetHandler)
		// http://blog.csdn.net/xsdxs/article/details/72849796
	    router.POST("/map", Routers.MappingHandler)
	}
