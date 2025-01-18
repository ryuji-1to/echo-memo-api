package main

import (
	"echo-rest-api/db"
	"echo-rest-api/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func main() {
	dbConnect := db.SetupDB()
	defer fmt.Println("Successfully Migrated")
	defer closeDB(dbConnect)
	dbConnect.AutoMigrate(&model.User{}, &model.Memo{})
}

func closeDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
