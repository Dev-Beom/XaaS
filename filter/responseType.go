package filter

import (
	"github.com/labstack/echo"
	"net/http"
)

type ResponseType struct {
	Code      int
	Interface echo.Map
}

func GetOKResponseType(k string, v interface{}) ResponseType {
	return ResponseType{
		Code: http.StatusOK,
		Interface: echo.Map{
			"status": true,
			k:        v,
		},
	}
}

func GetErrResponseType(code int, err error) ResponseType {
	return ResponseType{
		Code: code,
		Interface: echo.Map{
			"status":  false,
			"message": err.Error(),
		},
	}
}
