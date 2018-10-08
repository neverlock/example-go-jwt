package routers

import (
	"api.jwt.auth/controllers"
	"api.jwt.auth/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}

func SetLoginPage(router *mux.Router) *mux.Router {
	router.Handle("/login",
		negroni.New(
			negroni.HandlerFunc(controllers.LoginPageController),
		)).Methods("GET")
	return router
}
