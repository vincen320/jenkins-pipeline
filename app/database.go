package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/vincen320/user-service/helper"
)

func NewConnection() *sql.DB {
	var DB_DRIVER = os.Getenv("DB_DRIVER")

	if DB_DRIVER == "mysql" {
		return connectMySQL()
	}

	return connectPostgres()
}

func connectMySQL() *sql.DB {
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
		fmt.Println("mysql-ENV FILE NOT LOADED")
	} else {
		fmt.Println("mysql-LOAD FROM ENV FILE")
	}

	db, err := sql.Open("mysql", dbSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func connectPostgres() *sql.DB {
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_PORT = os.Getenv("DB_PORT")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_SSL_MODE = os.Getenv("DB_SSL_MODE")

	dbSourceName := "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + "/" + DB_NAME + "?sslmode=" + DB_SSL_MODE

	if DB_USER == "" || DB_HOST == "" || DB_PORT == "" || DB_NAME == "" {
		fmt.Println(DB_USER)
		fmt.Println(DB_PASSWORD)
		fmt.Println(DB_HOST)
		fmt.Println(DB_PORT)
		fmt.Println(DB_NAME)

		dbSourceName = "postgres://pqgouser:password@localhost/dbname?sslmode=disable"
		fmt.Println("postgres-ENV FILE NOT LOADED")
	} else {
		fmt.Println("postgres-LOAD FROM ENV FILE")
	}

	db, err := sql.Open("postgres", dbSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
