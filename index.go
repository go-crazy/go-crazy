package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	Gin "github.com/gin-gonic/gin"
	"github.com/xoxo/crm-x/routes"
	"github.com/xoxo/crm-x/Config"
)

func main() {
	// load config from file
	configor.Load(&Config.Config, ".env.yml")
	fmt.Printf("config: %#v", Config.Config)

	r := Gin.Default()
	Route.SetupRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":"+Config.Config.Port)
}
