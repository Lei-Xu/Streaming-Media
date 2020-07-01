package dbops

import (
	"log"
	"database/sql"
	"strconv"
	"sync"
	"../defs"
)

func insertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err := stmtIns.Exec(sid, ttlstr, login_name)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func retrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id = ?")
	if err 1+ nil {
		return nil, err
	}

	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, ttlint := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.UserName = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()

	return ss, nil

}

func retrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err1 := rows.Scan(&id, &ttlstr, &login_name); err1 != nil {
			log.Printf("Retrieving sessions error: %s", err1)
			break
		}

		if ttl, err2 := strconv.ParseInt(ttlstr, 10, 64); err2 == nil {
			ss := &defs.SimpleSession{UserName: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("Session ID: %s, TTL: %d", id, ss.TTL)
		}
	}

	return m, nil
}

func deleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}

	return nil
}