package myerror

import "net/http"

var (
	ErrGeneral     = New("something went wrong", http.StatusInternalServerError)
	ErrBodyRequest = New("failed get body request", http.StatusBadRequest)
)

type Error struct {
	Message    string
	StatusCode int
}

func New(msg string, statusCode int) Error {
	return Error{
		Message:    msg,
		StatusCode: statusCode,
	}
}

func (e Error) Error() string {
	return e.Message
}
