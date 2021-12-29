package filter

import (
	"errors"
	"github.com/labstack/echo"
	"net/http"
)

type ResponseType struct {
	Code      int
	Interface echo.Map
}

func GetOKResponseType(k string, v interface{}) (ResponseType, error) {
	return ResponseType{
		Code: http.StatusOK,
		Interface: echo.Map{
			"status": true,
			k:        v,
		},
	}, nil
}

func GetErrResponseType(code int, err error) (ResponseType, error) {
	return ResponseType{
		Code: code,
		Interface: echo.Map{
			"status":  false,
			"message": err.Error(),
		},
	}, errors.New(err.Error())
}
