package repository

import (
	"echo-rest-api/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IMemoRepository interface {
	GetAllMemos(memos *[]model.Memo, userId uint) error
	GetMemoById(memo *model.Memo, userId uint, memoId uint) error
	CreateMemo(memo *model.Memo) error
	UpdateMemo(memo *model.Memo, userId uint, memoId uint) error
	DeleteMemo(userId uint, memoId uint) error
}

type memoRepository struct {
	db *gorm.DB
}

func NewMemoRepository(db *gorm.DB) IMemoRepository {
	return &memoRepository{db}
}

func (mr *memoRepository) GetAllMemos(memos *[]model.Memo, userId uint) error {
	if err := mr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(memos).Error; err != nil {
		return err
	}
	return nil
}

func (mr *memoRepository) GetMemoById(memo *model.Memo, userId uint, memoId uint) error {
	if err := mr.db.Joins("User").Where("user_id = ?", userId).First(memo, memo).Error; err != nil {
		return err
	}
	return nil
}

func (mr *memoRepository) CreateMemo(memo *model.Memo) error {
	if err := mr.db.Create(memo).Error; err != nil {
		return err
	}
	return nil
}

func (mr *memoRepository) UpdateMemo(memo *model.Memo, userId uint, memoId uint) error {
	result := mr.db.Model(memo).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", memoId, userId).Update("title", memo.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (mr *memoRepository) DeleteMemo(userId uint, memoId uint) error {
	result := mr.db.Where("id = ? AND user_id = ?", memoId, userId).Delete(&model.Memo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
