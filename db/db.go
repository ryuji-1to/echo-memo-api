package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	env := os.Getenv("GO_ENV")

	if env == "test" {
		db, err = gorm.Open(sqlite.Open(":memory"), &gorm.Config{})
		fmt.Println("sqlite db")
	} else {
		if env == "dev" {
			err = godotenv.Load()
			if err != nil {
				log.Fatalln(err)
			}
		}
		url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PW"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"))
		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		fmt.Println("connected db")
	}

	if err != nil {
		log.Fatalln(err)
	}
	return db

}
