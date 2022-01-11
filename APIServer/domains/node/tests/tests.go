package tests

import "github.com/dev-beom/xaas/apiserver/models"

var (
	mockDB                = make(map[string]models.Node)
	nodeCreateRequestBody = models.Node{
		Id:          "test_id",
		Description: "test_description",
	}
)
