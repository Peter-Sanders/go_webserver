package db

import (
	"database/sql"
	"fmt"
	"log"
  "os"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db *sql.DB
}


func NewStore(dbName string) (Store, error) {
	Db, err := getConnection(dbName)
	if err != nil {
		return Store{}, err
	}

	return Store{
		Db,
	}, nil
}


func getConnection(dbName string) (*sql.DB, error) {
  db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("🔥 failed to connect to the database: %s", err)
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}


func Get_sql(file string) (string){
  b, err := os.ReadFile("sql/" + file + ".sql")
  if err != nil {
    return " "
  }
  str := string(b)
  return str
}

