package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/vincen320/user-service/helper"
)

func NewConnection() *sql.DB {
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_PORT = os.Getenv("DB_PORT")
	var DB_NAME = os.Getenv("DB_NAME")

	dbSourceName := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ")/" + DB_NAME

	if DB_USER == "" || DB_HOST == "" || DB_PORT == "" || DB_NAME == "" {
		fmt.Println(DB_USER)
		fmt.Println(DB_PASSWORD)
		fmt.Println(DB_HOST)
		fmt.Println(DB_PORT)
		fmt.Println(DB_NAME)
		dbSourceName = "root:@tcp(localhost:3306)/v_user"
		fmt.Println("ENV FILE NOT LOADED")
	} else {
		fmt.Println("LOAD FROM ENV FILE")
	}

	db, err := sql.Open("mysql", dbSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
