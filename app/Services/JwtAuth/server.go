package main

import (
	"github.com/xoxo/crm-x/app/Services/JwtAuth/routers"
	"github.com/xoxo/crm-x/app/Services/JwtAuth/settings"
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
