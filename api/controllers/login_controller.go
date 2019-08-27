package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/victorsteven/fullstack/api/auth"
	"github.com/victorsteven/fullstack/api/models"
	"github.com/victorsteven/fullstack/api/responses"
	"github.com/victorsteven/fullstack/api/utils/channels"
	"github.com/victorsteven/fullstack/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {
	user := models.User{}
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)

		err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = models.VerifyPassword(user.Password, password)
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return auth.CreateToken(user.ID)
	}
	return "", err
}
