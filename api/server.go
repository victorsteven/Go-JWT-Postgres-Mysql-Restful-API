package api

import (
	"fmt"
	"log"
	"os"

	// "github.com/victorsteven/fullstack/api/router"
	"github.com/joho/godotenv"
	"github.com/victorsteven/fullstack/api/controllers"
)

var server = controllers.Server{}

// func refreshTable() error {
// 	err := server.DB.Debug().DropTableIfExists(&models.User{}).Error
// 	if err != nil {
// 		return err
// 	}
// 	err = server.DB.Debug().AutoMigrate(&models.User{}).Error
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Successfully refreshed table")
// 	return nil
// }

func Run() {

	// config.Load()
	// seed.Load()
	// fmt.Println("this is after seeding fresh")
	fmt.Printf("running... at port %d\n\n", 8080)
	// listen(8080)

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// refreshTable()

	server.Run(":8080")

}
