package main

import (
	"github.com/go-crazy/go-crazy/app/Services/JwtAuth/routers"
	"github.com/go-crazy/go-crazy/app/Services/JwtAuth/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
