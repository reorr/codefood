package db

import (
	"codefood/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?parseTime=true"

	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString error...!!!")
	}

	err = db.Ping()

	if err != nil {
		panic("DSN Invalid : " + connectionString)
	}
}

func CreateCon() *sql.DB {
	return db
}