package api

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/victorsteven/fullstack/api/controllers"
	"github.com/victorsteven/fullstack/api/seed"
)

var server = controllers.Server{}

var (
	PORT       = 0
	SECRETKEY  []byte
	DBURL      = ""
	DBDRIVER   = ""
	DBUSER     = ""
	DBPASSWORD = ""
	DBNAME     = ""
)

func Load() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 8000
	}
	DBHOST := os.Getenv("DB_HOST")
	if DBHOST == "" {
		DBHOST = "127.0.0.1"
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	DBUSER = os.Getenv("DB_USER")
	DBPASSWORD = os.Getenv("DB_PASSWORD")
	DBNAME = os.Getenv("DB_NAME")
	SECRETKEY = []byte(os.Getenv("API_SECRET"))

}

func Run() {

	Load()

	// config.Load()

	server.Initialize(DBDRIVER, DBUSER, DBPASSWORD, DBNAME)

	seed.Load(server.DB)

	server.Run(":8080")

}
