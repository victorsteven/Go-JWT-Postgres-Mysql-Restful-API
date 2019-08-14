package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT      = 0
	SECRETKEY []byte
	DBURL     = ""
	DBDRIVER  = ""
)

func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		PORT = 8000
	}

	DBHOST := os.Getenv("DATABASE_HOST")
	if DBHOST == "" {
		DBHOST = "127.0.0.1"
	}
	DBDRIVER = os.Getenv("DB_DRIVER")

	DBURL = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), DBHOST, os.Getenv("DB_NAME"))

	SECRETKEY = []byte(os.Getenv("API_SECRET"))

}
