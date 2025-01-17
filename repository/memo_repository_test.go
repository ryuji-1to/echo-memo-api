package repository

import (
	"echo-rest-api/model"
	"echo-rest-api/testHelpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllMemos(t *testing.T) {
	db := testHelpers.SetupTestData()

	repository := NewMemoRepository(db)
	result := []model.Memo{}
	const userId = uint(1)
	err := repository.GetAllMemos(&result, userId)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(result))
}

func TestGetMemoById(t *testing.T) {
	db := testHelpers.SetupTestData()

	repository := NewMemoRepository(db)
	result := model.Memo{}
	const (
		userId = uint(1)
		memoId = uint(1)
	)
	err := repository.GetMemoById(&result, userId, memoId)
	assert.Equal(t, nil, err)
	assert.Equal(t, userId, (result.UserId))
	assert.Equal(t, memoId, (result.ID))
}

func TestCreateMemo(t *testing.T) {
	db := testHelpers.SetupTestData()

	repository := NewMemoRepository(db)
	const (
		userId = uint(2)
		memoId = uint(4)
	)
	input := model.Memo{
		Title:   "created",
		Content: "created memo",
		UserId:  userId,
	}
	err := repository.CreateMemo(&input)
	assert.Equal(t, nil, err)
	createdMemo := model.Memo{}
	err = repository.GetMemoById(&createdMemo, userId, memoId)
	assert.Equal(t, nil, err)
	assert.Equal(t, memoId, createdMemo.ID)
	assert.Equal(t, userId, createdMemo.UserId)
}

func TestUpdateMemo(t *testing.T) {
	db := testHelpers.SetupTestData()

	repository := NewMemoRepository(db)
	updateMemo := model.Memo{
		Title:   "updated memo1 title",
		Content: "updated memo1 content",
	}
	const (
		userId = uint(1)
		memoId = uint(1)
	)
	err := repository.UpdateMemo(&updateMemo, userId, memoId)
	assert.Equal(t, nil, err)
	updatedMemo := model.Memo{}
	err = repository.GetMemoById(&updatedMemo, userId, memoId)
	assert.Equal(t, nil, err)
	assert.Equal(t, memoId, updatedMemo.ID)
	assert.Equal(t, updateMemo.Title, updatedMemo.Title)
	assert.Equal(t, updateMemo.Content, updatedMemo.Content)
}

func TestDeleteMemo(t *testing.T) {
	db := testHelpers.SetupTestData()

	repository := NewMemoRepository(db)
	const (
		userId = uint(1)
		memoId = uint(1)
	)
	err := repository.DeleteMemo(userId, memoId)
	assert.Equal(t, nil, err)

	err = repository.DeleteMemo(userId, memoId)
	assert.Equal(t, "object does not exist", err.Error())
}
