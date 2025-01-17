package testHelpers

import (
	"echo-rest-api/db"
	"echo-rest-api/model"

	"gorm.io/gorm"
)

func SetupTestData() *gorm.DB {
	db := db.SetupDB()
	if db.Migrator().HasTable(&model.Memo{}) {
		db.Migrator().DropTable(&model.Memo{})
	}
	if db.Migrator().HasTable(&model.User{}) {
		db.Migrator().DropTable(&model.User{})
	}
	db.AutoMigrate(&model.Memo{}, &model.User{})
	memos := []model.Memo{
		{Title: "memo1 title", Content: "memo1 content", UserId: 1},
		{Title: "memo2 title", Content: "memo2 content", UserId: 2},
		{Title: "memo3 title", Content: "memo3 content", UserId: 1},
	}
	users := []model.User{
		{Email: "testuser1@example.com", Password: "testuser1"},
		{Email: "testuser2@example.com", Password: "testuser2"},
		{Email: "testuser3@example.com", Password: "testuser3"},
	}
	for _, item := range memos {
		db.Create(&item)
	}
	for _, item := range users {
		db.Create(&item)
	}
	return db
}
