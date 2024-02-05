package dbFunc

//ref doc: https://gosamples.dev/sqlite-intro/

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
  
)

//set filename
const fileName = "/home/runner/moregolearning/dbFunc/foo.db"

func CreateDB() {
	fmt.Println("Creating DB")

  // open database/sql using sqlite3 as driver and fileName as data source
  // return db object and err
  // if err is not nil, then log the error
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
  // defer close until actions have been run
	defer db.Close()

  //prep simple sql query to create table
	sqlStmt := `
  create table foo (id integer not null primary key, name text);
  delete from foo;
  `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
