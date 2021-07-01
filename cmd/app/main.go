package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/vlasove/test03/internal/app/models"
	"github.com/vlasove/test03/internal/app/storage"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env files found")
	}
}

func main() {
	config := storage.NewConfig()
	store := storage.New(config)

	log.Println("trying to open connection to database...")
	if err := store.Open(); err != nil {
		log.Fatal("fail while building connection:", err)
	}
	defer store.Close()
	log.Println("database connection successfuly opened")

	{
		// normally test.do_something()
		id := 2
		attr := &models.Attributes{
			FirstAttribute:  "test first",
			SecondAttribute: "test second",
		}
		timestamp := time.Now()
		if err := store.Repo().DoSomething(id, attr, timestamp); err != nil {
			log.Println(err)
		}
	}

	{
		// code 1 for test.get_db_error()
		if err := store.Repo().GetDBError(1); err != nil {
			log.Println(err)
		}
	}

	{
		// code 2 for test.get_db_error()
		if err := store.Repo().GetDBError(2); err != nil {
			log.Println(err)
		}
	}

	{
		// code 15 for test.get_db_error()
		if err := store.Repo().GetDBError(15); err != nil {
			log.Println(err)
		}
	}

}
