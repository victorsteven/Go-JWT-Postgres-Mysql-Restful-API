package api

import (
	"fmt"
	"log"
	"os"

	// "github.com/victorsteven/fullstack/api/router"
	"github.com/joho/godotenv"
	"github.com/victorsteven/fullstack/api/controllers"
)

func Run() {
	// config.Load()
	// seed.Load()
	// fmt.Println("this is after seeding fresh")
	fmt.Printf("running... at port %d\n\n", 8080)
	// listen(8080)

	server := controllers.Server{}

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	server.Run(":8080")

}
