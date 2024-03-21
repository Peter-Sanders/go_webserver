package db

import (
	"database/sql"
	"fmt"
	"log"

  util "github.com/Peter-Sanders/go_webserver/util"

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

	if err := createMigrations(Db); err != nil {
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


func createMigrations(db *sql.DB) error {
  stmt := util.Get_sql("create_users")

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
