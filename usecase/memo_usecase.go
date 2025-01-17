package usecase

import (
	"echo-rest-api/model"
	"echo-rest-api/repository"
	"echo-rest-api/validator"
)

type IMemoUsecase interface {
	GetAllMemos(userId uint) ([]model.MemoResponse, error)
	GetMemoById(userId uint, memoId uint) (model.MemoResponse, error)
	CreateMemo(memo model.Memo) (model.MemoResponse, error)
	UpdateMemo(memo model.Memo, userId uint, memoId uint) (model.MemoResponse, error)
	DeleteMemo(userId uint, memoId uint) error
}

type memoUsecase struct {
	mr repository.IMemoRepository
	mv validator.IMemoValidator
}

func NewMemoUsecase(mr repository.IMemoRepository, mv validator.IMemoValidator) IMemoUsecase {
	return &memoUsecase{mr, mv}
}

func (mu *memoUsecase) GetAllMemos(userId uint) ([]model.MemoResponse, error) {
	memos := []model.Memo{}
	if err := mu.mr.GetAllMemos(&memos, userId); err != nil {
		return nil, err
	}
	resMemos := []model.MemoResponse{}
	for _, memo := range memos {
		t := model.MemoResponse{
			ID:        memo.ID,
			Title:     memo.Title,
			Content:   memo.Content,
			CreatedAt: memo.CreatedAt,
			UpdatedAt: memo.UpdatedAt,
		}
		resMemos = append(resMemos, t)
	}
	return resMemos, nil
}

func (mu *memoUsecase) GetMemoById(userId uint, memoId uint) (model.MemoResponse, error) {
	memo := model.Memo{}
	if err := mu.mr.GetMemoById(&memo, userId, memoId); err != nil {
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

func (mu *memoUsecase) CreateMemo(memo model.Memo) (model.MemoResponse, error) {
	if err := mu.mv.MemoValidate(memo); err != nil {
		return model.MemoResponse{}, err
	}
	if err := mu.mr.CreateMemo(&memo); err != nil {
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

func (mu *memoUsecase) UpdateMemo(memo model.Memo, userId uint, memoId uint) (model.MemoResponse, error) {
	if err := mu.mv.MemoValidate(memo); err != nil {
		return model.MemoResponse{}, err
	}
	if err := mu.mr.UpdateMemo(&memo, userId, memoId); err != nil {
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

func (mu *memoUsecase) DeleteMemo(userId uint, memoId uint) error {
	if err := mu.mr.DeleteMemo(userId, memoId); err != nil {
		return err
	}
	return nil
}
