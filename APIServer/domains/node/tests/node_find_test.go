package tests

import (
	"github.com/dev-beom/xaas/apiserver/domains/node"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testInitialization(e *echo.Echo, method string, target string, body io.Reader) (*httptest.ResponseRecorder, echo.Context) {
	request := httptest.NewRequest(method, target, body)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	return recorder, context
}

func Test_Node_Find_API(t *testing.T) {
	app := echo.New()
	initMockNodeDataForTest()
	t.Run("(정상) 정상적으로 개별 노드의 정보를 확인할 수 있다.", func(t *testing.T) {
		id := getRandomID()
		recorder, context := testInitialization(app, http.MethodGet, "/api/node", nil)
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
		recorder, context := testInitialization(app, http.MethodGet, "/api/node", nil)
		context.SetPath("/:id")
		context.SetParamNames("id")
		context.SetParamValues(id)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Get(context)) {
			assert.Equal(t, http.StatusNotFound, recorder.Code)
		}
	})
	t.Run("(정상) 정상적으로 모든 노드의 정보를 확인할 수 있다.", func(t *testing.T) {
		recorder, context := testInitialization(app, http.MethodGet, "/api/nodes", nil)
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.GetAll(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
}
