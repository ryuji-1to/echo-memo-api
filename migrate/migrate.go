package main

import (
	"echo-rest-api/db"
	"echo-rest-api/model"
	"fmt"
)

func main() {
	dbConnect := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConnect)
	dbConnect.AutoMigrate(&model.User{}, &model.Memo{})
}
