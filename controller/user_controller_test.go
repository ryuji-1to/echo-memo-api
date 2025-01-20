package controller

import (
	"bytes"
	"echo-rest-api/model"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUp(t *testing.T) {
	input := model.User{
		Email:    "testsignup@example.com",
		Password: "testsignup",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	usecase := newMockUserUsecase()
	mockResponse := model.UserResponse{
		ID:    1,
		Email: input.Email,
	}
	usecase.(*mockUserUsecase).
		On("SignUp", mock.Anything).
		Return(mockResponse, nil)
	controller := NewUserController(usecase)
	controller.SignUp(mockContext)
	assert.Equal(t, http.StatusCreated, rec.Code)

	userJSON, err := json.Marshal(mockResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, string(userJSON), rec.Body.String())
	usecase.(*mockUserUsecase).AssertExpectations(t)
}

func TestSignUp_Error(t *testing.T) {
	input := model.User{
		Email:    "testsignup@example.com",
		Password: "testsignup",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	usecase := newMockUserUsecase()
	usecase.(*mockUserUsecase).
		On("SignUp", mock.Anything).
		Return(nil, errors.New("error"))
	controller := NewUserController(usecase)
	controller.SignUp(mockContext)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "\"error\"\n", rec.Body.String())
	usecase.(*mockUserUsecase).AssertExpectations(t)
}

func TestSignUp_BadRequest(t *testing.T) {
	// headerを設定しないので400エラーになる
	input := model.User{
		Email:    "testsignup@example.com",
		Password: "testsignup",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(inputJSON))
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	usecase := newMockUserUsecase()
	usecase.(*mockUserUsecase).
		On("SignUp", mock.Anything).
		Return(nil, errors.New("error"))
	controller := NewUserController(usecase)
	controller.SignUp(mockContext)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	usecase.(*mockUserUsecase).AssertNotCalled(t, "SignUp")

}

func TestLogin(t *testing.T) {
	input := model.User{
		Email:    "testlogin@example.com",
		Password: "testlogin",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	usecase := newMockUserUsecase()
	const token = "testToken"
	usecase.(*mockUserUsecase).
		On("Login", mock.Anything).
		Return(token, nil)
	controller := NewUserController(usecase)
	controller.Login(mockContext)
	assert.Equal(t, http.StatusOK, rec.Code)
	cookie := rec.Header().Get("Set-Cookie")
	assert.Contains(t, cookie, "token="+token)
	usecase.(*mockUserUsecase).AssertExpectations(t)
}

func TestLogin_Error(t *testing.T) {
	input := model.User{
		Email:    "testlogin@example.com",
		Password: "testlogin",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	usecase := newMockUserUsecase()
	usecase.(*mockUserUsecase).
		On("Login", mock.Anything).
		Return(nil, errors.New("error"))
	controller := NewUserController(usecase)
	controller.Login(mockContext)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "\"error\"\n", rec.Body.String())
	usecase.(*mockUserUsecase).AssertExpectations(t)
}

func TestLogin_BadRequest(t *testing.T) {
	// headerを設定しないので400エラー
	input := model.User{
		Email:    "testlogin@example.com",
		Password: "testlogin",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(inputJSON))
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	usecase := newMockUserUsecase()
	usecase.(*mockUserUsecase).
		On("Login", mock.Anything).
		Return(nil, errors.New("error"))
	controller := NewUserController(usecase)
	controller.Login(mockContext)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	usecase.(*mockUserUsecase).AssertNotCalled(t, "Login")
}

func TestLogout(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	userController := NewUserController(nil)
	userController.Logout(mockContext)
	assert.Equal(t, http.StatusOK, rec.Code)
	token := rec.Header().Get("Set-Cookie")
	assert.Contains(t, token, "token=")
}

func TestCsrfToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/csrf", nil)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	controller := NewUserController(nil)
	controller.CsrfToken(mockContext)
	assert.Equal(t, http.StatusOK, rec.Code)
	csrf, err := json.Marshal(echo.Map{"csrf_token": "test_csrf_token"})
	assert.Nil(t, err)
	assert.JSONEq(t, string(csrf), rec.Body.String())
}
