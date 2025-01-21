package repository

import (
	"echo-rest-api/db"
	"echo-rest-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByEmail(t *testing.T) {
	db := db.SetupDB()
	repository := NewUserRepository(db)
	user := model.User{}
	const email = "testuser1@example.com"
	err := repository.GetUserByEmail(&user, email)
	assert.Nil(t, err)
	assert.Equal(t, email, user.Email)
}

func TestCreateUser(t *testing.T) {
	db := db.SetupDB()
	repository := NewUserRepository(db)
	input := model.User{
		Email:    "createuser@example.com",
		Password: "createuser",
	}

	err := repository.CreateUser(&input)
	assert.Nil(t, err)

	createdUser := model.User{}
	err = repository.GetUserByEmail(&createdUser, input.Email)
	assert.Nil(t, err)
	assert.Equal(t, input.Email, createdUser.Email)
}
