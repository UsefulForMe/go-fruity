package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/mocks/services"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func setup() *gin.Engine {
	os.Setenv("ENV", "test")
	config.InitConfig()
	router := gin.Default()
	return router

}

func Test_login_should_return_status_code_200(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	router := setup()

	mockService := services.NewMockUserService(c)
	mockFirebaseService := services.NewMockFirebaseAuthService(c)

	mockService.EXPECT().Login(gomock.Any()).Return(&dto.LoginUserResponse{}, nil)

	uh := NewAuthHandler(
		mockService,
		mockFirebaseService,
	)

	router.POST("/v1/login", uh.Login())
	var jsonStr = []byte(`{"phone_number":"0338613062","id_token":"123456"}`)
	req, _ := http.NewRequest(http.MethodPost, "/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != 200 {
		t.Errorf("Expected status code to be 200, got %d", recorder.Code)
	}

}
