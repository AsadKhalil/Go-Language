package database

import (
	"RestAPI/pkg/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Importing the postgres driver
)

var db *sql.DB

func InitDb(config *config.Config) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	fmt.Print(config.DbName)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)

	}
}

func GetDb() *sql.DB {
	return db
}
