package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/victorsteven/fullstack/api/models"
	"gopkg.in/go-playground/assert.v1"
)

var postInstance = models.Post{}

func refreshUserAndPostTable() error {

	err := server.DB.Debug().DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	err = server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {

	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}
	user := models.User{
		Nickname: "Sam Phil",
		Email:    "sam@gmail.com",
		Password: "password",
	}
	err = server.DB.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}
	post := models.Post{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.ID,
	}
	err = server.DB.Debug().Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func seedUsersAndPosts() error {

	var err error

	if err != nil {
		return err
	}
	var users = []models.User{
		models.User{
			Nickname: "Steven victor",
			Email:    "steven@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Magu Frank",
			Email:    "magu@gmail.com",
			Password: "password",
		},
	}
	var posts = []models.Post{
		models.Post{
			Title:   "Title 1",
			Content: "Hello world 1",
		},
		models.Post{
			Title:   "Title 2",
			Content: "Hello world 2",
		},
	}

	for i, _ := range users {
		err = server.DB.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return nil
}

func TestCreatePost(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}

	samples := []struct {
		inputJSON    string
		statusCode   int
		title        string
		content      string
		author_id    int
		errorMessage string
	}{
		{
			inputJSON:    `{"title":"The title", "content": "the content", "author_id": 1}`,
			statusCode:   201,
			title:        "The title",
			content:      "the content",
			author_id:    1,
			errorMessage: "",
		},
		// {
		// 	inputJSON:    `{"title":"The title", "content": "the content", "author_id": 1}`,
		// 	statusCode:   500,
		// 	errorMessage: "Title Already Taken",
		// },
		// {
		// 	inputJSON:    `{"title": "", "content": "The content", "author_id": 1}`,
		// 	statusCode:   422,
		// 	errorMessage: "Required Title",
		// },
		// {
		// 	inputJSON:    `{"title": "This is a title", "content": "", "author_id": 1}`,
		// 	statusCode:   422,
		// 	errorMessage: "Required Content",
		// },
		// {
		// 	inputJSON:    `{"title": "This is an awesome title", "content": "the content", "author_id": ""}`,
		// 	statusCode:   422,
		// 	errorMessage: "Required Author",
		// },
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/posts", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			log.Fatalf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateUser)

		userTokenString := fmt.Sprintf("Bearer %v", token)
		req.Header.Set("Authorization", userTokenString)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["title"], v.title)
			assert.Equal(t, responseMap["content"], v.content)
			assert.Equal(t, responseMap["author_id"], v.author_id)
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

// func TestGetPosts(t *testing.T) {

// 	refreshPostTable()

// 	seedPosts()

// 	req, err := http.NewRequest("GET", "/users", nil)
// 	if err != nil {
// 		log.Fatalf("this is the error: %v\n", err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(server.GetUsers)
// 	handler.ServeHTTP(rr, req)

// 	// defer server.DB.Close()

// 	var users []models.User
// 	err = json.Unmarshal([]byte(rr.Body.String()), &users)

// 	assert.Equal(t, rr.Code, http.StatusOK)
// 	assert.Equal(t, len(users), 2)
// }

// func TestGetPostByID(t *testing.T) {

// 	refreshPostTable()

// 	user, err := seedOnePost()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	userSample := []struct {
// 		id           string
// 		statusCode   int
// 		nickname     string
// 		email        string
// 		errorMessage string
// 	}{
// 		{
// 			id:         strconv.Itoa(int(user.ID)),
// 			statusCode: 200,
// 			nickname:   user.Nickname,
// 			email:      user.Email,
// 		},
// 		{
// 			id:         "unknwon",
// 			statusCode: 400,
// 		},
// 	}

// 	for _, v := range userSample {

// 		req, _ := http.NewRequest("GET", "/users", nil)
// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})

// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.GetUser)
// 		handler.ServeHTTP(rr, req)

// 		responseMap := make(map[string]interface{})
// 		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 		if err != nil {
// 			fmt.Printf("Cannot convert to json: %v", err)
// 		}
// 		assert.Equal(t, rr.Code, v.statusCode)

// 		if v.statusCode == 200 {
// 			assert.Equal(t, user.Nickname, responseMap["nickname"])
// 			assert.Equal(t, user.Email, responseMap["email"])
// 		}
// 	}
// }

// func TestUpdatePost(t *testing.T) {

// 	refreshPostTable()

// 	seedPosts() //we need atleast two users to properly check the update

// 	samples := []struct {
// 		id           string
// 		updateJSON   string
// 		statusCode   int
// 		nickname     string
// 		email        string
// 		errorMessage string
// 	}{
// 		{
// 			id:           strconv.Itoa(int(1)),
// 			updateJSON:   `{"nickname":"Grand", "email": "grand@gmail.com", "password": "password"}`,
// 			statusCode:   200,
// 			nickname:     "Grand",
// 			email:        "grand@gmail.com",
// 			errorMessage: "",
// 		},
// 		{
// 			id:           strconv.Itoa(int(1)),
// 			updateJSON:   `{"nickname":"Frank", "email": "kenny@gmail.com", "password": "password"}`,
// 			statusCode:   500,
// 			errorMessage: "Email Already Taken",
// 		},
// 		{
// 			id:           strconv.Itoa(int(1)),
// 			updateJSON:   `{"nickname":"Kenny Morris", "email": "grand@gmail.com", "password": "password"}`,
// 			statusCode:   500,
// 			errorMessage: "Nickname Already Taken",
// 		},
// 		{
// 			id:           strconv.Itoa(int(1)),
// 			updateJSON:   `{"nickname":"Kan", "email": "kangmail.com", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Invalid Email",
// 		},
// 		{
// 			id:           strconv.Itoa(int(1)),
// 			updateJSON:   `{"nickname": "", "email": "kan@gmail.com", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Required Nickname",
// 		},
// 		{
// 			id:           strconv.Itoa(int(1)),
// 			updateJSON:   `{"nickname": "Kan", "email": "", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Required Email",
// 		},
// 		{
// 			id:         "unknwon",
// 			statusCode: 400,
// 		},
// 		{
// 			id:         strconv.Itoa(int(2)),
// 			updateJSON: `{"nickname": "Mike", "email": "mike@gmail.com", "password": "password"}`,
// 			statusCode: 401,
// 		},
// 	}

// 	for _, v := range samples {

// 		req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(v.updateJSON))

// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})

// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.UpdateUser)

// 		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1NjY2ODk5MjEsInVzZXJfaWQiOjF9.4CIgrndIbgUQELh6N2-y-w-pPTJRsHIoxnNY-izg7Kc")

// 		handler.ServeHTTP(rr, req)

// 		responseMap := make(map[string]interface{})
// 		err := json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 		if err != nil {
// 			fmt.Printf("Cannot convert to json: %v", err)
// 		}
// 		assert.Equal(t, rr.Code, v.statusCode)
// 		if v.statusCode == 200 {
// 			assert.Equal(t, responseMap["nickname"], v.nickname)
// 			assert.Equal(t, responseMap["email"], v.email)
// 		}
// 		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
// 			assert.Equal(t, responseMap["error"], v.errorMessage)
// 		}
// 	}
// }

// func TestDeletePost(t *testing.T) {

// 	refreshPostTable()

// 	seedPosts()

// 	userSample := []struct {
// 		id           string
// 		statusCode   int
// 		nickname     string
// 		email        string
// 		errorMessage string
// 	}{
// 		{
// 			id:         strconv.Itoa(int(1)),
// 			statusCode: 204,
// 		},
// 		{
// 			id:         "unknwon",
// 			statusCode: 400,
// 		},
// 		{
// 			id:         strconv.Itoa(int(2)),
// 			statusCode: 401,
// 		},
// 	}

// 	for _, v := range userSample {

// 		req, _ := http.NewRequest("GET", "/users", nil)
// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})

// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.DeleteUser)

// 		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1NjY2ODk5MjEsInVzZXJfaWQiOjF9.4CIgrndIbgUQELh6N2-y-w-pPTJRsHIoxnNY-izg7Kc")

// 		handler.ServeHTTP(rr, req)

// 		assert.Equal(t, rr.Code, v.statusCode)
// 	}
// }
