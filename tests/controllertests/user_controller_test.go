package controllertests

import (
	"log"

	"github.com/victorsteven/fullstack/api/models"
)

// defer server.DB.Close()

var userInstance = models.User{}

var AuthEmail, AuthNickname, AuthPassword string
var AuthID uint32

func refreshUserTable() error {
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

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Nickname: "Pet",
		Email:    "pet@gmail.com",
		Password: "password",
	}

	err = server.DB.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() ([]models.User, error) {

	var err error
	if err != nil {
		return nil, err
	}
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
			return []models.User{}, err
		}
	}
	return users, nil
}

// func TestCreateUser(t *testing.T) {

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	samples := []struct {
// 		inputJSON    string
// 		statusCode   int
// 		nickname     string
// 		email        string
// 		errorMessage string
// 	}{
// 		{
// 			inputJSON:    `{"nickname":"Pet", "email": "pet@gmail.com", "password": "password"}`,
// 			statusCode:   201,
// 			nickname:     "Pet",
// 			email:        "pet@gmail.com",
// 			errorMessage: "",
// 		},
// 		{
// 			inputJSON:    `{"nickname":"Frank", "email": "pet@gmail.com", "password": "password"}`,
// 			statusCode:   500,
// 			errorMessage: "Email Already Taken",
// 		},
// 		{
// 			inputJSON:    `{"nickname":"Pet", "email": "grand@gmail.com", "password": "password"}`,
// 			statusCode:   500,
// 			errorMessage: "Nickname Already Taken",
// 		},
// 		{
// 			inputJSON:    `{"nickname":"Kan", "email": "kangmail.com", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Invalid Email",
// 		},
// 		{
// 			inputJSON:    `{"nickname": "", "email": "kan@gmail.com", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Required Nickname",
// 		},
// 		{
// 			inputJSON:    `{"nickname": "Kan", "email": "", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Required Email",
// 		},
// 		{
// 			inputJSON:    `{"nickname": "Kan", "email": "kan@gmail.com", "password": ""}`,
// 			statusCode:   422,
// 			errorMessage: "Required Password",
// 		},
// 	}

// 	for _, v := range samples {

// 		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.inputJSON))
// 		if err != nil {
// 			t.Errorf("this is the error: %v", err)
// 		}
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.CreateUser)
// 		handler.ServeHTTP(rr, req)

// 		responseMap := make(map[string]interface{})
// 		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 		if err != nil {
// 			fmt.Printf("Cannot convert to json: %v", err)
// 		}
// 		assert.Equal(t, rr.Code, v.statusCode)
// 		if v.statusCode == 201 {
// 			assert.Equal(t, responseMap["nickname"], v.nickname)
// 			assert.Equal(t, responseMap["email"], v.email)
// 		}
// 		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
// 			assert.Equal(t, responseMap["error"], v.errorMessage)
// 		}
// 	}
// }

// func TestGetUsers(t *testing.T) {

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = seedUsers()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	req, err := http.NewRequest("GET", "/users", nil)
// 	if err != nil {
// 		t.Errorf("this is the error: %v\n", err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(server.GetUsers)
// 	handler.ServeHTTP(rr, req)

// 	// defer server.DB.Close()

// 	var users []models.User
// 	err = json.Unmarshal([]byte(rr.Body.String()), &users)
// 	if err != nil {
// 		log.Fatalf("Cannot convert to json: %v\n", err)
// 	}
// 	assert.Equal(t, rr.Code, http.StatusOK)
// 	assert.Equal(t, len(users), 2)
// }

// func TestGetUserByID(t *testing.T) {

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	user, err := seedOneUser()
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

// 		req, err := http.NewRequest("GET", "/users", nil)
// 		if err != nil {
// 			t.Errorf("This is the error: %v\n", err)
// 		}
// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.GetUser)
// 		handler.ServeHTTP(rr, req)

// 		responseMap := make(map[string]interface{})
// 		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 		if err != nil {
// 			log.Fatalf("Cannot convert to json: %v", err)
// 		}

// 		assert.Equal(t, rr.Code, v.statusCode)

// 		if v.statusCode == 200 {
// 			assert.Equal(t, user.Nickname, responseMap["nickname"])
// 			assert.Equal(t, user.Email, responseMap["email"])
// 		}
// 	}
// }

// func TestUpdateUser(t *testing.T) {

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	users, err := seedUsers() //we need atleast two users to properly check the update
// 	if err != nil {
// 		log.Fatalf("Error seeding user: %v\n", err)
// 	}
// 	// Get only the first user
// 	for _, user := range users {
// 		if user.ID == 2 {
// 			continue
// 		}
// 		AuthID = user.ID
// 		AuthNickname = user.Nickname
// 		AuthEmail = user.Email
// 		AuthPassword = user.Password
// 	}
// 	//Login the user and get the authentication token
// 	token, err := server.SignIn(AuthEmail, AuthPassword)
// 	if err != nil {
// 		log.Fatalf("cannot login: %v\n", err)
// 	}

// 	samples := []struct {
// 		id             string
// 		updateJSON     string
// 		statusCode     int
// 		updateNickname string
// 		updateEmail    string
// 		errorMessage   string
// 	}{
// 		{
// 			id:             strconv.Itoa(int(AuthID)),
// 			updateJSON:     `{"nickname":"Grand", "email": "grand@gmail.com", "password": "password"}`,
// 			statusCode:     200,
// 			updateNickname: "Grand",
// 			updateEmail:    "grand@gmail.com",
// 			errorMessage:   "",
// 		},
// 		{
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"nickname":"Frank", "email": "kenny@gmail.com", "password": "password"}`,
// 			statusCode:   500,
// 			errorMessage: "Email Already Taken",
// 		},
// 		{
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"nickname":"Kenny Morris", "email": "grand@gmail.com", "password": "password"}`,
// 			statusCode:   500,
// 			errorMessage: "Nickname Already Taken",
// 		},
// 		{
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"nickname":"Kan", "email": "kangmail.com", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Invalid Email",
// 		},
// 		{
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"nickname": "", "email": "kan@gmail.com", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Required Nickname",
// 		},
// 		{
// 			id:           strconv.Itoa(int(AuthID)),
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

// 		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.updateJSON))
// 		if err != nil {
// 			t.Errorf("This is the error: %v\n", err)
// 		}
// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})

// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.UpdateUser)

// 		userTokenString := fmt.Sprintf("Bearer %v", token)

// 		req.Header.Set("Authorization", userTokenString)

// 		handler.ServeHTTP(rr, req)

// 		responseMap := make(map[string]interface{})
// 		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 		if err != nil {
// 			fmt.Printf("Cannot convert to json: %v", err)
// 		}
// 		assert.Equal(t, rr.Code, v.statusCode)
// 		if v.statusCode == 200 {
// 			assert.Equal(t, responseMap["nickname"], v.updateNickname)
// 			assert.Equal(t, responseMap["email"], v.updateEmail)
// 		}
// 		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
// 			assert.Equal(t, responseMap["error"], v.errorMessage)
// 		}
// 	}
// }

// func TestDeleteUser(t *testing.T) {

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	users, err := seedUsers() //we need atleast two users to properly check the update
// 	if err != nil {
// 		log.Fatalf("Error seeding user: %v\n", err)
// 	}
// 	// Get only the first and log him in
// 	for _, user := range users {
// 		if user.ID == 2 {
// 			continue
// 		}
// 		AuthID = user.ID
// 		AuthNickname = user.Nickname
// 		AuthEmail = user.Email
// 		AuthPassword = user.Password
// 	}
// 	//Login the user and get the authentication token
// 	token, err := server.SignIn(AuthEmail, AuthPassword)
// 	if err != nil {
// 		log.Fatalf("cannot login: %v\n", err)
// 	}

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

// 		req, err := http.NewRequest("GET", "/users", nil)
// 		if err != nil {
// 			t.Errorf("This is the error: %v\n", err)
// 		}
// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.DeleteUser)

// 		userTokenString := fmt.Sprintf("Bearer %v", token)
// 		req.Header.Set("Authorization", userTokenString)

// 		handler.ServeHTTP(rr, req)
// 		assert.Equal(t, rr.Code, v.statusCode)
// 	}
// }
