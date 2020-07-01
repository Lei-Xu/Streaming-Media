package dbops

import (
	"database/sql"
	"time"
	"log"

	"../defs"
	"../utils"

	_ "github.com/go-sql-driver/mysql"
)

func addUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()

	return nil
}

func getUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()

	return pwd, nil
}

func deleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("Delete user error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	stmtDel.Close()

	return nil
}

func addNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.createUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("INSERT INTO videos (id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err !=- nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, name: name, DisplayCTime: ctime}

	defer stmtIns.Close()
	return res, nil
}


func getVideo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM videos WHERE id = ?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCTime: dct}
	retrurn res, nil
}

func deleteVideo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM videos WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func addNewComments(vid string, aid int, content string) error {
	id, err := utils.createUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err := stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()

	return nil
}

func listComments(vid  string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
	`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}

	defer stmtOut.Close()
	return res, nil
}