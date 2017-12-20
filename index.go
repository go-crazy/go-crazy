package main

import (
	"github.com/jinzhu/configor"
	Gin "github.com/gin-gonic/gin"
	"github.com/xoxo/crm-x/routes"
	. "github.com/xoxo/crm-x/Config"
)

func main() {
	// load config from file
	configor.Load(&Config, ".env.yml")
	// fmt.Printf("config: %#v\n\n\n", Config)

	// init path
	InitPath()

	// init logger
	InitLogger()

	// init database
	InitDB()

	// init gin engine
	r := Gin.Default()
	Route.SetupRouter(r)

	// Listen and Server in Config.Port
	r.Run(":"+Config.Port)
}
