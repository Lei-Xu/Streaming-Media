package defs

// Requests
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

// Data Model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCTime string
}

type Comment struct {
	Id       string
	VideoId  string
	Author int
	Content  string
}

type SimpleSession struct {
	UserName string
	TTL int64
}