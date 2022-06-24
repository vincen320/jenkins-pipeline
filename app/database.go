package app

import (
	"database/sql"
	"time"

	"github.com/vincen320/user-service/helper"
)

func NewConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/v_user")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
