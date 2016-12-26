package connector

import "fmt"

//Error  describe error message for connector
type Error struct {
	Code               int
	Message            string
	AdditionalErrorMsg string
}

const (
	//ErrCodeMethodNotAllow wechat request method not allow here
	ErrCodeMethodNotAllow int = iota + 1
	//ErrCodeInvalidBody invalid body
	ErrCodeInvalidBody
)

var errCodeToMessage = map[int]string{
	ErrCodeMethodNotAllow: "method not found",
}

func (e *Error) Error() string {
	return fmt.Sprintf("ConnectorError: %s, Code: %d, Message: %s, AdditionalErrorMsg: %s",
		errCodeToMessage[e.Code], e.Code, e.Message, e.AdditionalErrorMsg)
}

//Type indicate object type
func (e *Error) Type() string {
	return "error"
}

//NewMethodNotAllowError create a method not allow error
func NewMethodNotAllowError(message string, addition string) *Error {
	return &Error{
		Code:               ErrCodeMethodNotAllow,
		Message:            message,
		AdditionalErrorMsg: addition,
	}
}

//NewInvalidBodyError create invalid body error
func NewInvalidBodyError(message string, addition string) *Error {
	return &Error{
		Code:               ErrCodeInvalidBody,
		Message:            message,
		AdditionalErrorMsg: addition,
	}
}

// IsMethodNotAllow returns true if and only if err is "key" not found error.
func IsMethodNotAllow(err error) bool {
	return isErrCode(err, ErrCodeMethodNotAllow)
}

// IsInvalidBody returns true if and only if err is "key" not found error.
func IsInvalidBody(err error) bool {
	return isErrCode(err, ErrCodeInvalidBody)
}

func isErrCode(err error, code int) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*Error); ok {
		return e.Code == code
	}
	return false
}
