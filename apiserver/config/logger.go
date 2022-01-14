package config

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LoggerConfig() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}\t${time_rfc3339_nano}\t${uri}\t${status}\t${remote_ip}\tlatency=${latency}\n",
	})
}
