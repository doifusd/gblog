package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int                    `json:"code"`
	msg     string                 `json:"msg"`
	details []string               `json:"details"`
	data    map[string]interface{} `json:"data"`
}

var Codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := Codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d　已经存在，请更换一个", code))
	}
	Codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d ,错误信息:%s", e.Code(), e.Msg())
}
func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) Data() map[string]interface{} {
	return e.data
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

func (e *Error) WithData(data map[string]interface{}) *Error {
	newError := *e
	newError.data = data
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case IntvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequest.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
