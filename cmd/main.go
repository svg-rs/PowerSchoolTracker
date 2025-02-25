package main

import (
	"PowerSchoolTracker/pkg"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var err error
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var username string = os.Getenv("USERNAME1")
	var password string = os.Getenv("PASSWORD1")
	powerschooltracker.Track(username, password)
}
