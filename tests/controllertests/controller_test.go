package controllertests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/victorsteven/fullstack/api/controllers"
)

var server = controllers.Server{}

func TestMain(t *testing.T) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	TestDbHost := os.Getenv("TEST_DATABASE_HOST")

	if TestDbHost == "" {
		TestDbHost = "127.0.0.1"
	}
	TestDbDriver := os.Getenv("TestDbDriver")
	TestDbURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), TestDbHost, os.Getenv("TestDbName"))
	server.DB, err = gorm.Open(TestDbDriver, TestDbURL)
	if err != nil {
		log.Printf("Error connecting to the database: %v\n", err)
		return
	}
	// defer server.DB.Close()

	fmt.Println("We have connected to the database")
}
