package cmd

import (
	"RestAPI/api"
	"RestAPI/pkg/database"
	"log"
	"net/http"
)

func Run() {
	config := database.NewConfig()
	database.InitDb(config)

	r := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
	defer database.GetDb().Close()

}
