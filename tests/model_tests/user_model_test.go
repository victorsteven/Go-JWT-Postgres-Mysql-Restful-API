package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	// _ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/victorsteven/fullstack/api/controllers"
	"github.com/victorsteven/fullstack/api/models"
	"gopkg.in/go-playground/assert.v1"
)

var server = controllers.Server{}

var userInstance = models.User{}

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
	fmt.Println("We have connected to the database")
}

func refreshTable() error {
	err := server.DB.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() models.User {

	refreshTable()

	user := models.User{
		Nickname: "Kenny Morris",
		Email:    "kenny@gmail.com",
		Password: "password",
	}

	err := server.DB.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user
}

func seedUsers() {

	users := []models.User{
		models.User{
			Nickname: "Steven victor",
			Email:    "steven@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Kenny Morris",
			Email:    "kenny@gmail.com",
			Password: "password",
		},
	}

	for i, _ := range users {
		err := server.DB.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

func TestFindAllUsers(t *testing.T) {

	refreshTable()

	seedUsers()

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		fmt.Printf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {

	refreshTable()

	newUser := models.User{
		ID:       1,
		Email:    "test@gmail.com",
		Nickname: "test",
		Password: "password",
	}
	savedUser, err := newUser.SaveUser(server.DB)
	if err != nil {
		fmt.Printf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Nickname, savedUser.Nickname)
}

func TestGetUserByID(t *testing.T) {

	refreshTable()

	user := models.User{
		ID:       1,
		Nickname: "modi",
		Email:    "modi@gmail.com",
		Password: "password",
	}
	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		fmt.Printf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nickname, user.Nickname)
}

func TestUpdateAUser(t *testing.T) {

	refreshTable()

	user := seedOneUser()

	userUpdate := models.User{
		ID:       1,
		Nickname: "modiUpdate",
		Email:    "modiupdate@gmail.com",
		Password: "password",
	}
	updatedUser, err := userUpdate.UpdateAUser(server.DB, user.ID)
	if err != nil {
		fmt.Printf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Nickname, userUpdate.Nickname)

}

func TestDeleteAUser(t *testing.T) {

	refreshTable()

	user := seedOneUser()

	isDeleted, err := userInstance.DeleteAUser(server.DB, user.ID)
	if err != nil {
		fmt.Printf("this is the error updating the user: %v\n", err)
		return
	}
	// assert.Equal(t, int(isDeleted), 1) //one shows that the record has been deleted or:
	assert.Equal(t, isDeleted, int64(1)) //one shows that the record has been deleted
}
