package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type middlewareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddlewareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middlewareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func registerHandlers() *httprouter.Router {
	router := httprouter.New()

	router.Get("/videos/:vid-id", streamHandler)

	router.POST("/upload/:vid-id", uploadHandler)

	return router
}

func (m middlewareHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorresponse(w, http.StatusTooManyRequests, "Too many request")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := registerHandlers()
	mh := NewMiddlewareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}