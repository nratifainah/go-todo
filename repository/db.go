package repository

import (
	"database/sql"
	"fmt"
	"log"
	_ "os"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error

	dbUser := "root"
	dbName := "salary"
	dbPass := "root"
	dbHost := "127.0.0.1"
	dbPort := "3303"

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	log.Print(dbSource)
	dbConnect, err := sql.Open("mysql", dbSource)

	if err != nil {
		log.Fatal("DB ERROR")
		log.Fatal(err)
		panic(err)
	}

	log.Print("connected to DB")

	if err = dbConnect.Ping(); err != nil {
		log.Fatal("DB PING ERROR")
		log.Fatal(err)
	}

	db = dbConnect
}