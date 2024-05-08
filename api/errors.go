package api

import "net/http"

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

func (e Error) Error() string {
	return e.Err
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func InvalidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "Invalid ID",
	}
}

func ResourceNotFound() Error {
	return Error{
		Code: http.StatusNotFound,
		Err:  "Resource not found",
	}
}

func Unauthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "Unauthorized",
	}
}

func NotValidDatetime() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "Not valid datetime",
	}
}
