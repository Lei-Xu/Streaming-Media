package defs

type Error struct {
	ErrorMsg  string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Err    Error
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Err: Error{ErrorMsg: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{HttpSC: 401, Err: Error{ErrorMsg: "User authentication failed", ErrorCode: "002"}}
	ErrorDBError                = ErrorResponse{HttpSC: 500, Err: Error{ErrorMsg: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults = ErrorResponse{HttpSC: 500, Err: Error{ErrorMsg: "Internal service error", ErrorCode: "004"}}
)
