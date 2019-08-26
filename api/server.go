package api

import (
	"github.com/victorsteven/fullstack/api/controllers"
	"github.com/victorsteven/fullstack/api/seed"
	"github.com/victorsteven/fullstack/config"
)

var server = controllers.Server{}

func Run() {

	config.Load()

	server.Initialize(config.DBDRIVER, config.DBUSER, config.DBPASSWORD, config.DBNAME)

	seed.Load(server.DB)

	server.Run(":8080")

}
