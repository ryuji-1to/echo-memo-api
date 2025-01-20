package controller

import (
	"echo-rest-api/model"
	"echo-rest-api/usecase"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

func createMockContext(req *http.Request, rec *httptest.ResponseRecorder) echo.Context {
	e := echo.New()
	mockContext := e.NewContext(req, rec)
	mockContext.Set("csrf", "test_csrf_token")
	mockContext.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
			"exp":     time.Now().Add(time.Hour).Unix(),
		},
	})
	return mockContext
}

type mockMemoUsecase struct {
	mock.Mock
}

func newMockMemoUsecase() usecase.IMemoUsecase {
	return &mockMemoUsecase{}
}

func (m *mockMemoUsecase) GetAllMemos(userId uint) ([]model.MemoResponse, error) {
	args := m.Called(userId)
	if memoArg, ok := args.Get(0).([]model.MemoResponse); ok && memoArg != nil {
		return memoArg, nil
	}
	return nil, args.Error(1)
}

func (m *mockMemoUsecase) GetMemoById(userId uint, memoId uint) (model.MemoResponse, error) {
	args := m.Called(userId, memoId)
	if memoArg, ok := args.Get(0).(model.MemoResponse); ok {
		return memoArg, nil
	}
	return model.MemoResponse{}, args.Error(1)
}

func (m *mockMemoUsecase) CreateMemo(memo model.Memo) (model.MemoResponse, error) {
	args := m.Called(memo)
	if err, ok := args.Get(0).(error); ok && err != nil {
		return model.MemoResponse{}, err
	}
	resMemo := model.MemoResponse{
		ID:        memo.ID,
		Title:     memo.Title,
		Content:   memo.Content,
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
	}
	return resMemo, nil
}

func (m *mockMemoUsecase) UpdateMemo(memo model.Memo, userId uint, memoId uint) (model.MemoResponse, error) {
	args := m.Called(memo, userId, memoId)
	if err, ok := args.Get(0).(error); ok && err != nil {
		return model.MemoResponse{}, err
	}
	resMemo := model.MemoResponse{
		ID:        memo.ID,
		Title:     memo.Title,
		Content:   memo.Content,
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
	}
	return resMemo, nil
}

func (m *mockMemoUsecase) DeleteMemo(userId uint, memoId uint) error {
	args := m.Called(userId, memoId)
	return args.Error(0)
}

type mockUserUsecase struct {
	mock.Mock
}

func newMockUserUsecase() usecase.IUserUsecase {
	return &mockUserUsecase{}
}

func (m *mockUserUsecase) Login(user model.User) (string, error) {
	args := m.Called(user)
	if tokenArg, ok := args.Get(0).(string); ok && tokenArg != "" {
		return tokenArg, nil
	}
	return "", args.Error(1)
}

func (m *mockUserUsecase) SignUp(user model.User) (model.UserResponse, error) {
	args := m.Called(user)
	if userArg, ok := args.Get(0).(model.UserResponse); ok {
		resUser := model.UserResponse{
			ID:    userArg.ID,
			Email: userArg.Email,
		}
		return resUser, nil
	}
	return model.UserResponse{}, args.Error(1)
}
