package main

import (
	"net/http"
	"../api/session"
)

var HEADER_FILED_SESSION = "X-Session-Id"
vae HEADER_FILED_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FILED_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FILED_UNAME, uname)
	return true
}

func validateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FILED_UNAME)
	if len(uname) == 0 {
		sendErrorResponse()
		return false
	}

	return true
}