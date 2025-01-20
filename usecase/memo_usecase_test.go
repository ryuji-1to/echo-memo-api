package usecase

import (
	"echo-rest-api/model"
	"echo-rest-api/validator"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetAllMemos(t *testing.T) {
	const userId = uint(1)
	expectedMemos := []model.Memo{
		{Title: "mock memo1 title", Content: "mock memo1 content", UserId: userId},
		{Title: "mock memo2 title", Content: "mock memo2 content", UserId: userId},
	}
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("GetAllMemos", mock.Anything, userId).Return(&expectedMemos, nil)

	usecase := NewMemoUsecase(mockRepository, nil)
	memos, err := usecase.GetAllMemos(userId)
	assert.Nil(t, err)
	assert.Equal(t, len(expectedMemos), len(memos))
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestGetAllMemos_Error(t *testing.T) {
	const userId = uint(1)
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("GetAllMemos", mock.Anything, userId).Return(nil, errors.New("error"))

	usecase := NewMemoUsecase(mockRepository, nil)
	memos, err := usecase.GetAllMemos(userId)
	assert.Error(t, err)
	assert.Nil(t, memos)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestGetMemoById(t *testing.T) {
	const (
		userId = uint(1)
		memoId = uint(1)
	)
	expectedMemo := model.Memo{
		Model: gorm.Model{
			ID: memoId,
		},
		Title:   "mock memo1 title",
		Content: "mock memo1 content",
		UserId:  userId,
	}
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("GetMemoById", mock.Anything, userId).Return(&expectedMemo, nil)

	usecase := NewMemoUsecase(mockRepository, nil)
	memo, err := usecase.GetMemoById(userId, memoId)
	assert.Nil(t, err)
	assert.Equal(t, expectedMemo.ID, memo.ID)
	assert.Equal(t, expectedMemo.Title, memo.Title)
	assert.Equal(t, expectedMemo.Content, memo.Content)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestGetMemoById_Error(t *testing.T) {
	const (
		userId = uint(1)
		memoId = uint(1)
	)
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("GetMemoById", mock.Anything, userId).Return(nil, errors.New("error"))

	usecase := NewMemoUsecase(mockRepository, nil)
	memo, err := usecase.GetMemoById(userId, memoId)
	assert.Equal(t, model.MemoResponse{}, memo)
	assert.Error(t, err)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestCreateMemo(t *testing.T) {
	mockMemo := model.Memo{
		Title:   "mock memo1 title",
		Content: "mock memo1 content",
	}
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("CreateMemo", mock.Anything).Return(&mockMemo, nil)

	validator := validator.NewMemoValidator()
	usecase := NewMemoUsecase(mockRepository, validator)
	memo, err := usecase.CreateMemo(mockMemo)
	assert.Nil(t, err)
	assert.Equal(t, mockMemo.Title, memo.Title)
	assert.Equal(t, mockMemo.Content, memo.Content)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestCreateMemo_Error(t *testing.T) {
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("CreateMemo", mock.Anything).Return(nil, errors.New("error"))
	mockMemo := model.Memo{
		Title:   "mock memo1 title",
		Content: "mock memo1 content",
	}

	validator := validator.NewMemoValidator()
	usecase := NewMemoUsecase(mockRepository, validator)
	memo, err := usecase.CreateMemo(mockMemo)
	assert.Equal(t, model.MemoResponse{}, memo)
	assert.Error(t, err)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestCreateMemo_Validate(t *testing.T) {
	validator := validator.NewMemoValidator()
	usecase := NewMemoUsecase(nil, validator)
	mockMemo1 := model.Memo{
		Title: "",
	}
	memo, err := usecase.CreateMemo(mockMemo1)
	assert.Equal(t, "title: title is required.", err.Error())
	assert.Equal(t, model.MemoResponse{}, memo)

	mockMemo2 := model.Memo{
		Title: "Too long title should be validated. Too long title should be validated. Too long title should be validated.",
	}
	memo, err = usecase.CreateMemo(mockMemo2)
	assert.Equal(t, "title: limited max 50 length.", err.Error())
	assert.Equal(t, model.MemoResponse{}, memo)
}

func TestUpdateMemo(t *testing.T) {
	const (
		userId = uint(1)
		memoId = uint(1)
	)
	mockMemo := model.Memo{
		Model: gorm.Model{
			ID: memoId,
		},
		Title:   "updated mock memo1 title",
		Content: "updated mock memo1 content",
		UserId:  userId,
	}
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("UpdateMemo", mock.Anything, userId, memoId).Return(&mockMemo, nil)

	validator := validator.NewMemoValidator()
	usecase := NewMemoUsecase(mockRepository, validator)
	memo, err := usecase.UpdateMemo(mockMemo, userId, memoId)
	assert.Nil(t, err)
	assert.Equal(t, mockMemo.Title, memo.Title)
	assert.Equal(t, mockMemo.Content, memo.Content)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestUpdateMemo_Error(t *testing.T) {
	const (
		userId = uint(1)
		memoId = uint(1)
	)
	mockMemo := model.Memo{
		Model: gorm.Model{
			ID: memoId,
		},
		Title:   "updated mock memo1 title",
		Content: "updated mock memo1 content",
		UserId:  userId,
	}
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("UpdateMemo", mock.Anything, userId, memoId).Return(nil, errors.New("error"))

	validator := validator.NewMemoValidator()
	usecase := NewMemoUsecase(mockRepository, validator)
	memo, err := usecase.UpdateMemo(mockMemo, userId, memoId)
	assert.Equal(t, model.MemoResponse{}, memo)
	assert.Error(t, err)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestUpdateMemo_Validate(t *testing.T) {
	validator := validator.NewMemoValidator()
	usecase := NewMemoUsecase(nil, validator)

	mockMemo1 := model.Memo{
		Title: "",
	}
	memo, err := usecase.CreateMemo(mockMemo1)
	assert.Equal(t, "title: title is required.", err.Error())
	assert.Equal(t, model.MemoResponse{}, memo)

	mockMemo2 := model.Memo{
		Title: "Too long title should be validated. Too long title should be validated. Too long title should be validated.",
	}
	memo, err = usecase.CreateMemo(mockMemo2)
	assert.Equal(t, "title: limited max 50 length.", err.Error())
	assert.Equal(t, model.MemoResponse{}, memo)
}

func TestDeleteMemo(t *testing.T) {
	const (
		userId = uint(1)
		memoId = uint(1)
	)
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("DeleteMemo", userId, memoId).Return(nil)
	usecase := NewMemoUsecase(mockRepository, nil)

	err := usecase.DeleteMemo(userId, memoId)
	assert.Nil(t, err)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}

func TestDeleteMemo_Error(t *testing.T) {
	mockRepository := newMockMemoRepository()
	mockRepository.(*mockMemoRepository).On("DeleteMemo", uint(1), uint(1)).Return(errors.New("error"))

	usecase := NewMemoUsecase(mockRepository, nil)
	err := usecase.DeleteMemo(1, 1)
	assert.Error(t, err)
	mockRepository.(*mockMemoRepository).AssertExpectations(t)
}
