package controller

import (
	"bytes"
	"echo-rest-api/model"
	"echo-rest-api/usecase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllMemos(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/memos", nil)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	memoResponse := []model.MemoResponse{
		{
			ID:      1,
			Title:   "memo1 title",
			Content: "memo1 content",
		},
		{
			ID:      2,
			Title:   "memo2 title",
			Content: "memo2 content",
		},
	}
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("GetAllMemos", uint(1)).
		Return(memoResponse, nil)
	controller := NewMemoController(mockUsecase)

	err := controller.GetAllMemos(mockContext)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, mockContext.Response().Status)

	memoJSON, err := json.Marshal(memoResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, string(memoJSON), rec.Body.String())
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestGetAllMemos_Error(t *testing.T) {}

func TestGetMemoById(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/memos/1", nil)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockContext.SetPath("/memos/:memoId")
	mockContext.SetParamNames("memoId")
	mockContext.SetParamValues("1")
	memoResponse := model.MemoResponse{
		ID:      1,
		Title:   "memo1 title",
		Content: "memo1 content",
	}
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("GetMemoById", uint(1), uint(1)).
		Return(memoResponse, nil)
	controller := NewMemoController(mockUsecase)

	err := controller.GetMemoById(mockContext)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, mockContext.Response().Status)

	memoJSON, err := json.Marshal(memoResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, string(memoJSON), rec.Body.String())
}

func TestGetMemoById_Error(t *testing.T) {}

func TestCreateMemo(t *testing.T) {
	input := model.Memo{
		Title:   "created memo",
		Content: "created memo",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/memos/1", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("CreateMemo", mock.Anything).
		Return(nil)
	controller := NewMemoController(mockUsecase)
	err = controller.CreateMemo(mockContext)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, mockContext.Response().Status)

	memoJSON, err := json.Marshal(model.MemoResponse{
		ID:        input.ID,
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	})
	assert.Nil(t, err)
	assert.JSONEq(t, string(memoJSON), rec.Body.String())
}

func TestCreateMemo_Error(t *testing.T) {}

func TestUpdateMemo(t *testing.T) {
	input := model.Memo{
		Title:   "updated memo",
		Content: "updated memo",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPut, "/memos/1", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockContext.SetPath("/memos/:memoId")
	mockContext.SetParamNames("memoId")
	mockContext.SetParamValues("1")
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("UpdateMemo", mock.AnythingOfType("model.Memo"), uint(1), uint(1)).
		Return(nil)
	controller := NewMemoController(mockUsecase)
	err = controller.UpdateMemo(mockContext)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, mockContext.Response().Status)

	memoJSON, err := json.Marshal(model.MemoResponse{
		ID:        input.ID,
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	})
	assert.Nil(t, err)
	assert.JSONEq(t, string(memoJSON), rec.Body.String())
}

func TestUpdateMemo_Error(t *testing.T) {}

func TestDeleteMemo(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/memos/1", nil)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockContext.SetPath("/memos/:memoId")
	mockContext.SetParamNames("memoId")
	mockContext.SetParamValues("1")
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("DeleteMemo", uint(1), uint(1)).
		Return(nil)
	controller := NewMemoController(mockUsecase)
	err := controller.DeleteMemo(mockContext)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, mockContext.Response().Status)
}

func TestDeleteMemo_Error(t *testing.T) {}

func createMockContext(req *http.Request, rec *httptest.ResponseRecorder) echo.Context {
	e := echo.New()
	mockContext := e.NewContext(req, rec)
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
