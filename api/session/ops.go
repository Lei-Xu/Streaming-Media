package session

import (
	"time"
	"sync"
	"../defs"
	"../dbops"
	"../utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.deleteSession(sid)
}

func loadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return 
	}
	r.range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func generateNewSessionId(un string) string {
	id, _ := utils.createUUID()
	ct := nowInMilli()
	ttl := ct + 30 * 60 * 1000

	ss := &def.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func isSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}