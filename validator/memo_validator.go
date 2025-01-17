package validator

import (
	"echo-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IMemoValidator interface {
	MemoValidate(memo model.Memo) error
}

type memoValidator struct{}

func NewMemoValidator() IMemoValidator {
	return &memoValidator{}
}

func (tv *memoValidator) MemoValidate(memo model.Memo) error {
	return validation.ValidateStruct(&memo,
		validation.Field(
			&memo.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 50).Error("limited max 50 length"),
		),
	)
}
