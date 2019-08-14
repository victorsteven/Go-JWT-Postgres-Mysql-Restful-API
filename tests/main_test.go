package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"github.com/victorsteven/fullstack/api/interfaces"
	"github.com/victorsteven/fullstack/api/models"
	"github.com/victorsteven/fullstack/api/services"
)

func TestMain(m *testing.M) {
	// err := testData.Unmarshal(&conn)
	// if err != nil {
	// 	log.Fatalf("unable to decode into struct, %v", err)
	// }

	// a = controllers.App{}
	// a.Initialize(conn.Host, conn.DbUser, conn.DbPass, conn.DbName)

	// ensureTableExists()
	// code := m.Run()
	// os.Exit(code)

	// a = services.DbInstance{}
	// a.

}

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	repository interfaces.UserInterface
	service    *services.DbInstance
	// user       *models.User
}

func (s *Suite) SetupSuite() {
	// var (
	// 	// SECRETKEY_TEST []byte
	// 	DBURL_TEST    = ""
	// 	DBDRIVER_TEST = ""
	// )
	var err error
	err = godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	}

	DBHOST_TEST := os.Getenv("DATABASE_HOST_TEST")
	if DBHOST_TEST == "" {
		DBHOST_TEST = "127.0.0.1"
	}

	DBDRIVER_TEST := os.Getenv("DB_DRIVER_TEST")
	DBURL_TEST := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASSWORD_TEST"), DBHOST_TEST, os.Getenv("DB_NAME_TEST"))
	s.DB, err = gorm.Open(DBDRIVER_TEST, DBURL_TEST)
	if err != nil {
		log.Printf("Error connecting to the database: %v", err)
	}
	// s.service = services.NewDbInstance(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

// func (r *DbInstance) SaveUser(user models.User) (models.User, error) {
// 	var err error
// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		defer close(ch)
// 		err = r.DB.Debug().Model(&models.User{}).Create(&user).Error
// 		if err != nil {
// 			ch <- false
// 			return
// 		}
// 		ch <- true
// 	}(done)

// 	if channels.OK(done) {
// 		return user, nil
// 	}
// 	return models.User{}, err
// }

func (s *Suite) TestSaveUser() {
	fmt.Println("this is the beginning of the save user test")
	// var (
	// 	ID       uint32 = 1
	// 	Nickname        = "steven"
	// 	Email           = "steven@gmail.com"
	// )
	user := models.User{
		ID:       1,
		Nickname: "haku",
		Email:    "steven@gmail.com",
		Password: "testing",
	}

	// fmt.Printf("Everything is good here: %v", user)
	res, err := s.service.SaveUser(user)

	if err != nil {
		log.Printf("this is the error creating the user: %v", err)
	}
	log.Printf("this is the created user: %v", res)

	// response, token := services.Login(&user)
	// assert.Equal(t, http.StatusOK, response)
	// assert.NotEmpty(t, token)
}

// func (s *Suite) TestFindAllUsers() {
// 	fmt.Println("This is the beginning of the test")
// 	var (
// 		ID       uint32 = 1
// 		Nickname        = "steven"
// 		Email           = "steven@gmail.com"
// 	)

// 	res, err := s.repository.FindUserByID(ID)

// 	log.Printf("thiis is the error calling the function: %v", err)
// 	fmt.Printf("this is the response: %v", res)
// }

// func (s *Suite) TestFindAllUsers() {
// 	// users := []struct {
// 	// 	ID       uint32
// 	// 	Nickname string
// 	// 	Email    string
// 	// }{
// 	// 	{
// 	// 		ID:       1,
// 	// 		Nickname: "steven",
// 	// 		Email:    "steven@gmail.com",
// 	// 	},
// 	// 	{
// 	// 		ID:       2,
// 	// 		Nickname: "mike",
// 	// 		Email:    "mike@gmail.com",
// 	// 	},
// 	// }
// 	// type users struct {
// 	// 	ID       uint32
// 	// 	Nickname string
// 	// 	Email    string
// 	// }
// 	// p2 := users{
// 	// 	ID:       1,
// 	// 	Nickname: "steven",
// 	// 	Email:    "steven@gmail.com",
// 	// }
// 	// fmt.Printf("This is the user: %v", p2)

// 	var (
// 		ID       uint32 = 1
// 		Nickname        = "steven"
// 		Email           = "steven@gmail.com"
// 	)
// 	s.mock.ExpectQuery("^SELECT (.+) FROM users  WHERE id = (.+)$").WithArgs(ID).WillReturnRows(sqlmock.NewRows([]string{"ID", "Nickname", "Email"}).AddRow(ID, Nickname, Email))

// 	// s.mock.ExpectQuery(regexp.QuoteMeta(
// 	// 	`SELECT * FROM "users" WHERE (id = $1)`)).WithArgs(ID).WillReturnRows(sqlmock.NewRows([]string{"ID", "Nickname", "Email"}).AddRow(ID, Nickname, Email))

// 	// This is a comment
// 	res, err := s.repository.FindUserByID(ID)
// 	require.NoError(s.T(), err)
// 	require.Nil(s.T(), deep.Equal(&models.User{ID: ID, Nickname: Nickname, Email: Email}, res))
// }
