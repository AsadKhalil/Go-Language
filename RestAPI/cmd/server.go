package cmd

import (
	"RestAPI/api"
	"RestAPI/pkg/config"
	"RestAPI/pkg/database"
	"log"
	"net/http"
)

func Run() {
	config := config.NewConfig()
	database.InitDb(config)

	r := api.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
	defer database.GetDb().Close()

}
