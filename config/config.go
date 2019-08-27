package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SECRETKEY  = []byte(os.Getenv("API_SECRET"))
	DBDRIVER   = os.Getenv("DB_DRIVER")
	DBUSER     = os.Getenv("DB_USER")
	DBPASSWORD = os.Getenv("DB_PASSWORD")
	DBNAME     = os.Getenv("DB_NAME")
	DBPORT     = os.Getenv("DB_PORT")
	DBHOST     = os.Getenv("DB_HOST")
)

func Load() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}
