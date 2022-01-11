package tests

import (
	"github.com/dev-beom/xaas/apiserver/domains/node"
	"github.com/dev-beom/xaas/apiserver/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Node_Delete_API(t *testing.T) {
	app := echo.New()
	t.Run("(정상) 정상적으로 노드가 삭제된다.", func(t *testing.T) {
		mockDB["test"] = models.Node{}
		request := httptest.NewRequest(http.MethodDelete, "/api/node", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		context.SetPath("/:id")
		context.SetParamNames("id")
		context.SetParamValues("test")
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Delete(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
		}
	})
	t.Run("(비정상) 해당 노드가 존재하지 않은 경우 오류가 발생한다.", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/api/node", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		context.SetPath("/:id")
		context.SetParamNames("id")
		context.SetParamValues("test")
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Delete(context)) {
			assert.Equal(t, http.StatusNotFound, recorder.Code)
		}
	})
	t.Run("(비정상) 노드 ID 없이 삭제할 경우 오류가 발생한다.", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/api/node", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		context := app.NewContext(request, recorder)
		context.SetPath("/:id")
		context.SetParamNames("id")
		context.SetParamValues("test")
		handler := node.NewController(node.NewService(node.NewMockRepository(mockDB)))
		if assert.NoError(t, handler.Delete(context)) {
			assert.Equal(t, http.StatusNotFound, recorder.Code)
		}
	})
}
