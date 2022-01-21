package node

import (
	"github.com/dev-beom/xaas/apiserver/filter"
	"github.com/dev-beom/xaas/apiserver/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

/**
todo 3 layer 구조 완성
- node 관련 end point 정리
- validation 코드 추가
*/

type controller struct {
	nodeService Service
	validate    *validator.Validate
}

func NewController(service Service) *controller {
	return &controller{
		service,
		validator.New(),
	}
}

// Get [GET] /api/node/:id
func (c *controller) Get(context echo.Context) error {
	id := context.Param("id")
	node, err := c.nodeService.Get(id)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusNotFound, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	resp := filter.GetOKResponseType("Data", node)
	return context.JSON(resp.Code, resp.Interface)
}

// GetAll [GET] /api/nodes
func (c *controller) GetAll(context echo.Context) error {
	nodes := c.nodeService.GetAll()
	resp := filter.GetOKResponseType("Data", nodes)
	return context.JSON(resp.Code, resp.Interface)
}

// Create [POST] /api/node
func (c *controller) Create(context echo.Context) error {
	nodeCreateRequestDto := new(models.NodeCreateRequestDto)
	err := context.Bind(nodeCreateRequestDto)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusInternalServerError, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	err = c.validate.Struct(nodeCreateRequestDto)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusBadRequest, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	var node = models.Node{
		Id:          nodeCreateRequestDto.Id,
		Description: nodeCreateRequestDto.Description,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	node.SetStateCreating()
	err = c.nodeService.Create(node)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusBadRequest, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	resp := filter.GetOKResponseType("Data", node)
	return context.JSON(resp.Code, resp.Interface)
}

// Delete [DELETE] /api/node/:id
func (c *controller) Delete(context echo.Context) error {
	id := context.Param("id")
	err := c.nodeService.Delete(id)
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
	updatedNode, err := c.nodeService.UpdateDescription(id, description)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusNotFound, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	resp := filter.GetOKResponseType("Data", updatedNode)
	return context.JSON(resp.Code, resp.Interface)
}

func (c *controller) FileUpload(context echo.Context) error {
	var nodeID string
	var newFileName string
	qParams := context.QueryParams()
	if val, ok := qParams["id"]; ok {
		nodeID = val[0]
	}
	if val, ok := qParams["name"]; ok {
		newFileName = val[0]
	}
	form, err := context.MultipartForm()
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusBadRequest, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	err = c.nodeService.FileUpload(nodeID, newFileName, form)
	if err != nil {
		resp := filter.GetErrResponseType(http.StatusBadRequest, err)
		return context.JSON(resp.Code, resp.Interface)
	}
	resp := filter.GetOKResponseType("Data", true)
	return context.JSON(resp.Code, resp.Interface)
}
