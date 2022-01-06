package instance

import (
	"github.com/dev-beom/faas/models"
	"github.com/dev-beom/faas/utils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockDB                      = make(map[string]models.Instance)
	instanceCreateRequestBodyOk = models.Instance{
		Id:          "test_id",
		Description: "test_description",
	}
)

func TestCreateInstance(t *testing.T) {
	app := echo.New()
	buffer, _ := utils.InterfaceToBuffer(instanceCreateRequestBodyOk)
	request := httptest.NewRequest(http.MethodPost, "/api/instance", &buffer)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recoder := httptest.NewRecorder()
	context := app.NewContext(request, recoder)
	handler := NewController(NewService(NewMockRepository(mockDB)))
	if assert.NoError(t, handler.Create(context)) {
		assert.Equal(t, http.StatusOK, recoder.Code)
	}
	utils.LoggingByTest("", recoder.Code, recoder.Body.String())
}
