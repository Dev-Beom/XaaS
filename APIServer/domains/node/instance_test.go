package node

import (
	"github.com/dev-beom/xaas/apiserver/models"
	"github.com/dev-beom/xaas/apiserver/utils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockDB                  = make(map[string]models.Node)
	nodeCreateRequestBodyOk = models.Node{
		Id:          "test_id",
		Description: "test_description",
	}
)

func Test_User_Create(t *testing.T) {
	t.Run("(정상) 정상적으로 노드가 생성된다", func(t *testing.T) {
		app := echo.New()
		//mockDB["test_id"] = models.Node{}
		buffer, _ := utils.InterfaceToBuffer(nodeCreateRequestBodyOk)
		request := httptest.NewRequest(http.MethodPost, "/api/node", &buffer)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recoder := httptest.NewRecorder()
		context := app.NewContext(request, recoder)
		handler := NewController(NewService(NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Create(context)) {
			assert.Equal(t, http.StatusOK, recoder.Code)
		}
		utils.LoggingByTest("", recoder.Code, recoder.Body.String())
	})
}
