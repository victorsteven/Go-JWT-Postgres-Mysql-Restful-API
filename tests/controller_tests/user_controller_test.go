package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
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
	// defer server.DB.Close()

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

func seedOneUser() (models.User, error) {

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
	return user, nil
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

type sampleCases struct {
	inputJson    string
	statusCode   int
	nickname     string
	email        string
	errorMessage string
}

func TestCreateUser(t *testing.T) {

	refreshTable()

	samples := []sampleCases{
		sampleCases{
			inputJson:    `{"nickname":"Pet", "email": "pet@gmail.com", "password": "password"}`,
			statusCode:   201,
			nickname:     "Pet",
			email:        "pet@gmail.com",
			errorMessage: "",
		},
		sampleCases{
			inputJson:    `{"nickname":"Frank", "email": "pet@gmail.com", "password": "password"}`,
			statusCode:   500,
			errorMessage: "Email Already Taken",
		},
		sampleCases{
			inputJson:    `{"nickname":"Pet", "email": "grand@gmail.com", "password": "password"}`,
			statusCode:   500,
			errorMessage: "Nickname Already Taken",
		},
		sampleCases{
			inputJson:    `{"nickname":"Kan", "email": "kangmail.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Invalid Email",
		},
		sampleCases{
			inputJson:    `{"nickname": "", "email": "kan@gmail.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required Nickname",
		},
		sampleCases{
			inputJson:    `{"nickname": "Kan", "email": "", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required Email",
		},
		sampleCases{
			inputJson:    `{"nickname": "Kan", "email": "kan@gmail.com", "password": ""}`,
			statusCode:   422,
			errorMessage: "Required Password",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.inputJson))
		if err != nil {
			log.Fatalf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateUser)
		handler.ServeHTTP(rr, req)

		fmt.Printf("This is the response payload: %v", rr.Body)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["nickname"], v.nickname)
			assert.Equal(t, responseMap["email"], v.email)
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestGetUsers(t *testing.T) {

	refreshTable()

	seedUsers()

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		log.Fatalf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUsers)
	handler.ServeHTTP(rr, req)

	// defer server.DB.Close()

	var users []models.User
	err = json.Unmarshal([]byte(rr.Body.String()), &users)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(users), 2)
}

func TestGetUserByID(t *testing.T) {

	refreshTable()

	user, err := seedOneUser()

	fmt.Printf("this is the user: %v", user)
	if err != nil {
		log.Fatal(err)
	}
	id := []byte(fmt.Sprintf(`{"id":%d}`, user.ID))
	req, err := http.NewRequest("GET", "/users", bytes.NewReader(id))
	if err != nil {
		log.Fatalf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUser)
	handler.ServeHTTP(rr, req)

	fmt.Printf("this is the awesome: %v", rr.Body)

	// responseMap := make(map[string]interface{})
	// err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
	// if err != nil {
	// 	fmt.Printf("Cannot convert to json: %v", err)
	// }

	// fmt.Printf("This is the response json %v", responseMap)

	// assert.Equal(t, rr.Code, http.StatusCreated)
	// assert.Equal(t, expectedMap["nickname"], responseMap["nickname"])
	// assert.Equal(t, responseMap["email"], responseMap["email"])

}

// func TestCreateUser(t *testing.T) {

// 	refreshTable()

// 	payload := `{"nickname":"Pet", "email": "pet@gmail.com", "password": "password"}`
// 	req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(payload))
// 	if err != nil {
// 		log.Fatalf("this is the error: %v", err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(server.CreateUser)
// 	handler.ServeHTTP(rr, req)

// 	defer server.DB.Close()

// 	if status := rr.Code; status != http.StatusCreated {
// 		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}
// 	jsonStr := `{"nickname":"Pet", "email": "pet@gmail.com", "password": "password"}`
// 	expectedMap := make(map[string]interface{})
// 	err = json.Unmarshal([]byte(jsonStr), &expectedMap)
// 	if err != nil {
// 		fmt.Printf("Cannot convert to json: %v", err)
// 	}
// 	responseMap := make(map[string]interface{})
// 	err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 	if err != nil {
// 		fmt.Printf("Cannot convert to json: %v", err)
// 	}
// 	assert.Equal(t, rr.Code, http.StatusCreated)
// 	assert.Equal(t, expectedMap["nickname"], responseMap["nickname"])
// 	assert.Equal(t, responseMap["email"], responseMap["email"])

// }
