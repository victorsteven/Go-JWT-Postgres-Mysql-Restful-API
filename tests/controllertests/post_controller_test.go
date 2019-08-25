package controllertests

// "strconv"
// "reflect"

// func refreshPostTable() error {
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

// func seedOnePost() (models.User, error) {

// 	refreshPostTable()

// 	user := models.User{
// 		Nickname: "Pet",
// 		Email:    "pet@gmail.com",
// 		Password: "password",
// 	}

// 	err := server.DB.Debug().Model(&models.User{}).Create(&user).Error
// 	if err != nil {
// 		log.Fatalf("cannot seed users table: %v", err)
// 	}
// 	return user, nil
// }

// func seedPosts() {

// 	users := []models.User{
// 		models.User{
// 			Nickname: "Steven victor",
// 			Email:    "steven@gmail.com",
// 			Password: "password",
// 		},
// 		models.User{
// 			Nickname: "Kenny Morris",
// 			Email:    "kenny@gmail.com",
// 			Password: "password",
// 		},
// 	}

// 	for i, _ := range users {
// 		err := server.DB.Debug().Model(&models.User{}).Create(&users[i]).Error
// 		if err != nil {
// 			log.Fatalf("cannot seed users table: %v", err)
// 		}
// 	}
// }

// func TestCreatePost(t *testing.T) {

// 	refreshPostTable()

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
// 			log.Fatalf("this is the error: %v", err)
// 		}
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.CreateUser)
// 		handler.ServeHTTP(rr, req)

// 		// fmt.Printf("This is the response payload: %v", rr.Body)

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
