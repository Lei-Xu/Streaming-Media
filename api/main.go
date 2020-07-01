package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middlewareHandler struct {
	r *httprouter.Router
}

func createMiddlewareHandler(r *httprouter.Router) http.Handler {
	m := middlewareHandler{}
	m.r = r
	return m
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// Sign up
	router.POST("/user", createUser)
	// Sign in
	router.POST("/user/:user_name", login)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := createMiddlewareHandler(r)
	http.ListenAndServe(":8000", r)
}
