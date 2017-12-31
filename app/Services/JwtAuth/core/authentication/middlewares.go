package authentication

import (
	"fmt"
	Gin "github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"net/http"
)

func RequireTokenAuthentication() Gin.HandlerFunc {
	return func(c *Gin.Context) {
		// before request
		authBackend := InitJWTAuthenticationBackend()
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			} else {
				return authBackend.PublicKey, nil
			}
		})

		//&& !authBackend.IsInBlacklist(c.Request.Header.Get("Authorization"))
		fmt.Println(token.Claims)
		if err == nil && token.Valid  {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
