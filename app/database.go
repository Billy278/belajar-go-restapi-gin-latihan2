package app

import (
	"belajar-go-restapi-gin-latihan2/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restapi_gin_latihan2")
	helper.PanicIfError(err)
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(20)
	//meamtikan konkesi  apabila tidak digunakan lagi
	DB.SetConnMaxIdleTime(10 * time.Minute)

	//refresh coneksi kembali ke minimal database
	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}
