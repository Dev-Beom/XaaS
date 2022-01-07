package instance

import (
	"github.com/dev-beom/xaas/filter"
	"github.com/dev-beom/xaas/models"
	"github.com/labstack/echo"
	"net/http"
	"time"
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

// Get [GET] /api/instance/:id
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

// GetAll [GET] /api/instances
func (c *controller) GetAll(context echo.Context) error {
	instances := c.instanceService.GetAll()
	resp := filter.GetOKResponseType("Data", instances)
	return context.JSON(resp.Code, resp.Interface)
}

// Create [POST] /api/instance
func (c *controller) Create(context echo.Context) error {
	var instanceCreateRequestDto models.InstanceCreateRequestDto
	err := context.Bind(&instanceCreateRequestDto)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusInternalServerError, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	var instance = models.Instance{
		Id:          instanceCreateRequestDto.Id,
		Description: instanceCreateRequestDto.Description,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	instance.SetStateCreating()
	err = c.instanceService.Create(instance)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusBadRequest, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	resp := filter.GetOKResponseType("Data", true)
	return context.JSON(resp.Code, resp.Interface)
}

// Delete [DELETE] /api/instance/:id
func (c *controller) Delete(context echo.Context) error {
	id := context.Param("id")
	err := c.instanceService.Delete(id)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusNotFound, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	resp := filter.GetOKResponseType("Data", true)
	return context.JSON(resp.Code, resp.Interface)
}

func (c *controller) UpdateDescription(context echo.Context) error {
	params := make(map[string]string)
	_ = context.Bind(&params)
	id := context.Param("id")
	description := params["description"]
	updatedInstance, err := c.instanceService.UpdateDescription(id, description)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusNotFound, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	resp := filter.GetOKResponseType("Data", updatedInstance)
	return context.JSON(resp.Code, resp.Interface)
}
