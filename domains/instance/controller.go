package instance

import (
	"github.com/dev-beom/faas/filter"
	"github.com/labstack/echo"
	"net/http"
)

/**
todo 3 layer 구조 완성
- instance 관련 end point 정리
- validation 코드 추가
*/

type controller struct {
	instanceService Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

func (c *controller) Get(context echo.Context) error {
	id := context.Param("id")
	instance, err := c.instanceService.Get(id)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusNotFound, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	resp := filter.GetOKResponseType("Data", instance)
	return context.JSON(resp.Code, resp.Interface)
}
