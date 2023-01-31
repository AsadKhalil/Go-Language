package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Importing the postgres driver
)

var db *sql.DB

func InitDb(config *Config) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DbPassword, config.DbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func GetDb() *sql.DB {
	return db
}
