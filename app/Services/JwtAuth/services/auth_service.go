package services

import (
	"github.com/go-crazy/go-crazy/app/Services/JwtAuth/parameters"
	"github.com/go-crazy/go-crazy/app/Services/JwtAuth/core/authentication"
	"github.com/go-crazy/go-crazy/app/Services/JwtAuth/services/models"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"net/http"
)

func Login(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.Id)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{token,authentication.TokenType,authentication.ExpireOffset})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.Id)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token,authentication.TokenType,authentication.ExpireOffset})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}
