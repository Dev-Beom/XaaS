package tests

import (
	"github.com/dev-beom/xaas/apiserver/domains/node"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Node_Find_API(t *testing.T) {
	app := echo.New()
	initMockNodeDataForTest()
	t.Run("(정상) 정상적으로 개별 노드의 정보를 확인할 수 있다.", func(t *testing.T) {
		id := getRandomID()
		request := httptest.NewRequest(http.MethodGet, "/api/node", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		context.SetPath("/:id")
		context.SetParamNames("id")
		context.SetParamValues(id)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Get(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Contains(t, recorder.Body.String(), getMockNode(id).Id, getMockNode(id).State, getMockNode(id).Description)
		}
	})
	t.Run("(비정상) 해당 노드가 존재하지 않은 경우 오류가 발생한다.", func(t *testing.T) {
		id := "test"
		request := httptest.NewRequest(http.MethodGet, "/api/node", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		context.SetPath("/:id")
		context.SetParamNames("id")
		context.SetParamValues(id)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Get(context)) {
			assert.Equal(t, http.StatusNotFound, recorder.Code)
		}
	})
	t.Run("(정상) 정상적으로 모든 노드의 정보를 확인할 수 있다.", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api/nodes", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.GetAll(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}
