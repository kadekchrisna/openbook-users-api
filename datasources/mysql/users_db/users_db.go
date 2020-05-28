package clientdbs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// Client db connection
	Client *sql.DB
)

func init() {
	var err error
	datasoucres := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("ADDRESS"),
		os.Getenv("DB"))
	Client, err = sql.Open("mysql", datasoucres)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfully connected to db.")

}
