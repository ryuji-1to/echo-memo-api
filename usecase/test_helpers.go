package usecase

import (
	"echo-rest-api/model"
	"echo-rest-api/repository"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type mockMemoRepository struct {
	mock.Mock
}

func newMockMemoRepository() repository.IMemoRepository {
	return &mockMemoRepository{}
}

func (m *mockMemoRepository) GetAllMemos(memos *[]model.Memo, userId uint) error {
	args := m.Called(memos, userId)
	if memoArg, ok := args.Get(0).(*[]model.Memo); ok && memoArg != nil {
		*memos = *memoArg
	}
	return args.Error(1)
}

func (m *mockMemoRepository) GetMemoById(memo *model.Memo, userId uint, memoId uint) error {
	args := m.Called(userId, memoId)
	if memoArg, ok := args.Get(0).(*model.Memo); ok && memoArg != nil {
		*memo = *memoArg
	}
	return args.Error(1)

}

func (m *mockMemoRepository) CreateMemo(memo *model.Memo) error {
	args := m.Called(memo)
	if memoArg, ok := args.Get(0).(*model.Memo); ok && memoArg != nil {
		*memo = *memoArg
	}
	return args.Error(1)
}

func (m *mockMemoRepository) UpdateMemo(memo *model.Memo, userId uint, memoId uint) error {
	args := m.Called(memo, userId, memoId)
	if memoArg, ok := args.Get(0).(*model.Memo); ok && memoArg != nil {
		*memo = *memoArg
	}
	return args.Error(1)
}

func (m *mockMemoRepository) DeleteMemo(userId uint, memoId uint) error {
	args := m.Called(userId, memoId)
	return args.Error(0)
}

type mockUserRepository struct {
	mock.Mock
}

func newMockUserRepository() repository.IUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) CreateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepository) GetUserByEmail(user *model.User, email string) error {
	args := m.Called(user, email)
	if userArg, ok := args.Get(0).(*model.User); ok && userArg != nil {
		*user = *userArg
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		(*user).Password = string(hashedPassword)
	}
	return args.Error(1)
}
