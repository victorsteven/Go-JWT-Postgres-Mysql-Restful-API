package router

import (
	"net/http"

	"github.com/victorsteven/fullstack/api/controllers"
)

var loginRoutes = []Route{
	Route{
		Uri:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
