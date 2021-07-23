package main

import (
	"log"

	"github.com/Ibra1994/go-start/migrations"
	"github.com/Ibra1994/go-start/server"
)

func main() {

	migrations.CreateDatabase()
	router := server.NewRouter()
	log.Print(router)

	router.Run()
}
