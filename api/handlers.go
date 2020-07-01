package main

import (
	"encoding/json"
	"goVideo/api/defs"
	"io"
	"io/ioutil"
	"net/http"

	"./defs"
	"github.com/julienschmidt/httprouter"
)

func createUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.addUserCrential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	id := session.generateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name")
	io.WriteString(w, userName)
}
