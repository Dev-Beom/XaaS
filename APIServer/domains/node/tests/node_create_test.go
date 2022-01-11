package tests

import (
	"github.com/dev-beom/xaas/apiserver/domains/node"
	"github.com/dev-beom/xaas/apiserver/models"
	"github.com/dev-beom/xaas/apiserver/utils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Node_Create(t *testing.T) {
	t.Run("(정상) 정상적으로 노드가 생성된다", func(t *testing.T) {
		app := echo.New()
		buffer, _ := utils.InterfaceToBuffer(nodeCreateRequestBody)
		request := httptest.NewRequest(http.MethodPost, "/api/node", &buffer)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Create(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
	t.Run("(비정상) 동일한 노드를 생성할 경우 오류가 발생한다.", func(t *testing.T) {
		app := echo.New()
		mockDB["test_id"] = models.Node{}
		buffer, _ := utils.InterfaceToBuffer(nodeCreateRequestBody)
		request := httptest.NewRequest(http.MethodPost, "/api/node", &buffer)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Create(context)) {
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		}
	})
	t.Run("(비정상) 노드의 정보를 입력하지 않고 생성할 경우 오류가 발생한다.", func(t *testing.T) {
		app := echo.New()
		buffer, _ := utils.InterfaceToBuffer(models.NodeCreateRequestDto{})
		request := httptest.NewRequest(http.MethodPost, "/api/node", &buffer)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Create(context)) {
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		}
	})
	t.Run("(비정상) 노드 ID의 길이를 2미만으로 생성할경우 오류가 발생한다.", func(t *testing.T) {
		app := echo.New()
		nodeCreateRequestBody.Id = "_"
		buffer, _ := utils.InterfaceToBuffer(nodeCreateRequestBody)
		request := httptest.NewRequest(http.MethodPost, "/api/node", &buffer)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Create(context)) {
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		}
	})
}
