package main

import (
	Gin "github.com/gin-gonic/gin"
	"time"
	"log"
)

var DB = make(map[string]string)

func Logger() Gin.HandlerFunc {
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

		// bw := responseWriter{buffer: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		
		c.JSON(c.Writer.Status(), Gin.H{"user": c.MustGet("api_response"), "Message": "hey", "Number": 123})

	}
}

func api_response(c *Gin.Context,value Gin.H)  {
	c.Set("api_response",value)
}

func setupRouter() *Gin.Engine {
	// Disable Console Color
	// Gin.DisableConsoleColor()
	r := Gin.Default()


	// Global middleware
	// Logger middleware will write the logs to Gin.DefaultWriter even you set with GIN_MODE=release.
	// By default Gin.DefaultWriter = os.Stdout
	r.Use(Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(Gin.Recovery())

	// Ping test
	r.GET("/ping", func(c *Gin.Context) {
		 c.String(200, "pong")
		 //c.Set("v",Gin.H{"user": "", "status": "no value"})
		 api_response(c,Gin.H{"user": "", "status": "no value"})
	})

	// Get user value
	r.GET("/user/:name", func(c *Gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, Gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, Gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses Gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(Gin.BasicAuth(Gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", Gin.BasicAuth(Gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *Gin.Context) {
		user := c.MustGet(Gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, Gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
