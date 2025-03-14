package request

import "fmt"

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (r *Error) Error() string {
	return fmt.Sprintf("Code %d: %s", r.Code, r.Message)
}
