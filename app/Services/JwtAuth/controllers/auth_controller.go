package controllers

import (
	"fmt"
	Gin "github.com/gin-gonic/gin"
	"github.com/xoxo/crm-x/app/Services/JwtAuth/services"
	"github.com/xoxo/crm-x/app/Services/JwtAuth/services/models"
	"encoding/json"
	"net/http"
)

func Login(c *Gin.Context) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(responseStatus)
	c.Writer.Write(token)
}

func RefreshToken(c *Gin.Context) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&requestUser)

	fmt.Println(requestUser)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(services.RefreshToken(requestUser))
}

func Logout(c *Gin.Context) {
	err := services.Logout(c.Request)
	c.Writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	} else {
		c.Writer.WriteHeader(http.StatusOK)
	}
}
