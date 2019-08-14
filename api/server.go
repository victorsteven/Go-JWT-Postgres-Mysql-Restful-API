package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/victorsteven/fullstack/api/router"
	"github.com/victorsteven/fullstack/api/seed"
	"github.com/victorsteven/fullstack/config"
)

func Run() {
	config.Load()
	seed.Load()
	fmt.Println("this is after seeding fresh")
	fmt.Printf("running... at port %d\n\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router.LoadCORS(r)))
}
