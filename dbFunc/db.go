package dbFunc

//ref doc: https://gosamples.dev/sqlite-intro/

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// set filename
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
  create table if not exists foo (id integer not null primary key, name text);
  delete from foo;
  `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func AddtoDB(name string) string {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	// defer close until actions have been run
	defer db.Close()
	sqlStmt := "insert into foo (name) values (?)"

	_, err = db.Exec(sqlStmt, name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return "error"
	}
	fmt.Println("ADDED TO DB")
	return "Added to DB"
}

func ReadDB() {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	// defer close until actions have been run
	defer db.Close()
	sqlStmt := "select * from foo"

	rows, err := db.Query(sqlStmt)
  if err != nil {
    fmt.Println("error")
  }

	defer db.Close()
  type Users struct {
      ID            int64  `field:"id"`                      
      Username      string `field:"name"`                              
  }
  for rows.Next() {
    user := new(Users)
    rows.Scan(&user.ID,&user.Username)
    fmt.Println("user: ", user)
    }
	return
}
