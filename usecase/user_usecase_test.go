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

func TestSignUp(t *testing.T) {
	mockUser := model.User{
		Email:    "testsignup@example.com",
		Password: "testsignup",
	}
	mockRepository := newMockUserRepository()
	mockRepository.(*mockUserRepository).On("CreateUser", mock.AnythingOfType("*model.User")).Return(nil)
	validator := validator.NewUserValidator()
	usecase := NewUserUsecase(mockRepository, validator)

	user, err := usecase.SignUp(mockUser)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, mockUser.Email)
	mockRepository.(*mockUserRepository).AssertExpectations(t)
}

func TestSignUp_Error(t *testing.T) {
	mockUser := model.User{
		Email:    "testsignup@example.com",
		Password: "testsignup",
	}
	mockRepository := newMockUserRepository()
	mockRepository.(*mockUserRepository).On("CreateUser", mock.AnythingOfType("*model.User")).Return(errors.New("error"))
	validator := validator.NewUserValidator()
	usecase := NewUserUsecase(mockRepository, validator)

	user, err := usecase.SignUp(mockUser)
	assert.Error(t, err)
	assert.Equal(t, model.UserResponse{}, user)
	mockRepository.(*mockUserRepository).AssertExpectations(t)
}

func TestSignUp_Validate(t *testing.T) {
	mockRepository := newMockUserRepository()
	mockRepository.(*mockUserRepository).On("CreateUser", mock.AnythingOfType("*model.User")).Return(errors.New("error"))
	validator := validator.NewUserValidator()
	usecase := NewUserUsecase(mockRepository, validator)

	mockUser := model.User{
		Email:    "",
		Password: "testsignup",
	}
	user, err := usecase.SignUp(mockUser)
	assert.Equal(t, "email: email is required.", err.Error())
	assert.Equal(t, model.UserResponse{}, user)

	mockUser = model.User{
		Email:    "thisistoolongemail@toolongemail.com",
		Password: "testsignup",
	}
	user, err = usecase.SignUp(mockUser)
	assert.Equal(t, "email: limited max 30 char.", err.Error())
	assert.Equal(t, model.UserResponse{}, user)

	mockUser = model.User{
		Email:    "testsignup",
		Password: "testsignup",
	}
	user, err = usecase.SignUp(mockUser)
	assert.Equal(t, "email: invalid email format.", err.Error())
	assert.Equal(t, model.UserResponse{}, user)

	mockUser = model.User{
		Email:    "testsignup@example.com",
		Password: "",
	}
	user, err = usecase.SignUp(mockUser)
	assert.Equal(t, "password: password is required.", err.Error())
	assert.Equal(t, model.UserResponse{}, user)

	mockUser = model.User{
		Email:    "testsignup@example.com",
		Password: "12345",
	}
	user, err = usecase.SignUp(mockUser)
	assert.Equal(t, "password: limited min 6 max 30 char.", err.Error())
	assert.Equal(t, model.UserResponse{}, user)
	mockRepository.(*mockUserRepository).AssertNotCalled(t, "CreateUser")
}

func TestLogin(t *testing.T) {
	mockUser := model.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email:    "testlogin@example.com",
		Password: "testlogin",
	}
	mockRepository := newMockUserRepository()
	mockRepository.(*mockUserRepository).On("GetUserByEmail", mock.AnythingOfType("*model.User"), mock.Anything).Return(&mockUser, nil)
	validator := validator.NewUserValidator()
	usecase := NewUserUsecase(mockRepository, validator)

	token, err := usecase.Login(mockUser)
	assert.NotEmpty(t, token)
	assert.Nil(t, err)
	mockRepository.(*mockUserRepository).AssertExpectations(t)
}

func TestLogin_Error(t *testing.T) {
	mockUser := model.User{
		Email:    "testlogin@example.com",
		Password: "testlogin",
	}
	mockRepository := newMockUserRepository()
	mockRepository.(*mockUserRepository).On("GetUserByEmail", mock.AnythingOfType("*model.User"), mock.Anything).Return(nil, errors.New("error"))

	validator := validator.NewUserValidator()
	usecase := NewUserUsecase(mockRepository, validator)
	token, err := usecase.Login(mockUser)
	assert.Empty(t, token)
	assert.Error(t, err)
	mockRepository.(*mockUserRepository).AssertExpectations(t)
}

func TestLogin_Validate(t *testing.T) {
	validator := validator.NewUserValidator()
	usecase := NewUserUsecase(nil, validator)

	mockUser := model.User{
		Email:    "",
		Password: "testsignup",
	}
	token, err := usecase.Login(mockUser)
	assert.Equal(t, "email: email is required.", err.Error())
	assert.Empty(t, token)

	mockUser = model.User{
		Email:    "thisistoolongemail@toolongemail.com",
		Password: "testsignup",
	}
	token, err = usecase.Login(mockUser)
	assert.Equal(t, "email: limited max 30 char.", err.Error())
	assert.Empty(t, token)

	mockUser = model.User{
		Email:    "testsignup",
		Password: "testsignup",
	}
	token, err = usecase.Login(mockUser)
	assert.Equal(t, "email: invalid email format.", err.Error())
	assert.Empty(t, token)

	mockUser = model.User{
		Email:    "testsignup@example.com",
		Password: "",
	}
	token, err = usecase.Login(mockUser)
	assert.Equal(t, "password: password is required.", err.Error())
	assert.Empty(t, token)

	mockUser = model.User{
		Email:    "testsignup@example.com",
		Password: "12345",
	}
	token, err = usecase.Login(mockUser)
	assert.Equal(t, "password: limited min 6 max 30 char.", err.Error())
	assert.Empty(t, token)
}
