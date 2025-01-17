package main

import (
	"echo-rest-api/controller"
	"echo-rest-api/db"
	"echo-rest-api/repository"
	"echo-rest-api/router"
	"echo-rest-api/usecase"
	"echo-rest-api/validator"
)

func main() {
	db := db.SetupDB()
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	memoRepository := repository.NewMemoRepository(db)
	memoValidator := validator.NewMemoValidator()
	memoUsecase := usecase.NewMemoUsecase(memoRepository, memoValidator)
	memoController := controller.NewMemoController(memoUsecase)
	e := router.NewRouter(userController, memoController)
	e.Logger.Fatal((e.Start(":8080")))
}
