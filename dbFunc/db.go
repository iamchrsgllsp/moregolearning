package dbFunc

//ref doc: https://gosamples.dev/sqlite-intro/

import (
	"database/sql"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"log"
	"path/filepath"
)

func CreateDB() {
	fmt.Println("Creating DB")
	fileName := "test.db"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Println(db)
}
