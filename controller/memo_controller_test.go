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

	controller.GetAllMemos(mockContext)
	assert.Equal(t, http.StatusOK, rec.Code)

	memoJSON, err := json.Marshal(memoResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, string(memoJSON), rec.Body.String())
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestGetAllMemos_Error(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/memos", nil)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("GetAllMemos", uint(1)).
		Return(nil, errors.New("error"))
	controller := NewMemoController(mockUsecase)

	controller.GetAllMemos(mockContext)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

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

	controller.GetMemoById(mockContext)
	assert.Equal(t, http.StatusOK, rec.Code)

	memoJSON, err := json.Marshal(memoResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, string(memoJSON), rec.Body.String())
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestGetMemoById_Error(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/memos/1", nil)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockContext.SetPath("/memos/:memoId")
	mockContext.SetParamNames("memoId")
	mockContext.SetParamValues("1")
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("GetMemoById", uint(1), uint(1)).
		Return(nil, errors.New("error"))
	controller := NewMemoController(mockUsecase)

	controller.GetMemoById(mockContext)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestCreateMemo(t *testing.T) {
	input := model.Memo{
		Title:   "created memo",
		Content: "created memo",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/memos", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("CreateMemo", mock.Anything).
		Return(nil)
	controller := NewMemoController(mockUsecase)

	controller.CreateMemo(mockContext)
	assert.Equal(t, http.StatusCreated, rec.Code)

	memoJSON, err := json.Marshal(model.MemoResponse{
		ID:        input.ID,
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	})
	assert.Nil(t, err)
	assert.JSONEq(t, string(memoJSON), rec.Body.String())
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestCreateMemo_Error(t *testing.T) {
	input := model.Memo{
		Title:   "created memo",
		Content: "created memo",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/memos", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("CreateMemo", mock.Anything).
		Return(errors.New("error"))
	controller := NewMemoController(mockUsecase)

	controller.CreateMemo(mockContext)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestCreateMemo_BadRequest(t *testing.T) {
	input := model.Memo{
		Title:   "created memo",
		Content: "created memo",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	// headerを付与しないので400になる
	req := httptest.NewRequest(http.MethodPost, "/memos", bytes.NewBuffer(inputJSON))
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("CreateMemo", mock.Anything).
		Return(errors.New("error"))
	controller := NewMemoController(mockUsecase)

	controller.CreateMemo(mockContext)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUsecase.(*mockMemoUsecase).AssertNotCalled(t, "CreateMemo")
}

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

	controller.UpdateMemo(mockContext)
	assert.Equal(t, http.StatusOK, rec.Code)

	memoJSON, err := json.Marshal(model.MemoResponse{
		ID:        input.ID,
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	})
	assert.Nil(t, err)
	assert.JSONEq(t, string(memoJSON), rec.Body.String())
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestUpdateMemo_Error(t *testing.T) {
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
		Return(errors.New("error"))
	controller := NewMemoController(mockUsecase)

	controller.UpdateMemo(mockContext)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestUpdateMemo_BadRequest(t *testing.T) {
	input := model.Memo{
		Title:   "updated memo",
		Content: "updated memo",
	}
	inputJSON, err := json.Marshal(input)
	assert.Nil(t, err)
	// headerを付与しないので400エラー
	req := httptest.NewRequest(http.MethodPut, "/memos/1", bytes.NewBuffer(inputJSON))
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockContext.SetPath("/memos/:memoId")
	mockContext.SetParamNames("memoId")
	mockContext.SetParamValues("1")
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("UpdateMemo", mock.AnythingOfType("model.Memo"), uint(1), uint(1)).
		Return(errors.New("error"))
	controller := NewMemoController(mockUsecase)

	controller.UpdateMemo(mockContext)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUsecase.(*mockMemoUsecase).AssertNotCalled(t, "UpdateMemo")
}

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

	controller.DeleteMemo(mockContext)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}

func TestDeleteMemo_Error(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/memos/1", nil)
	rec := httptest.NewRecorder()
	mockContext := createMockContext(req, rec)
	mockContext.SetPath("/memos/:memoId")
	mockContext.SetParamNames("memoId")
	mockContext.SetParamValues("1")
	mockUsecase := newMockMemoUsecase()
	mockUsecase.(*mockMemoUsecase).
		On("DeleteMemo", uint(1), uint(1)).
		Return(errors.New("error"))
	controller := NewMemoController(mockUsecase)

	controller.DeleteMemo(mockContext)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.(*mockMemoUsecase).AssertExpectations(t)
}
