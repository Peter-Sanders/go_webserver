package server

import (
   "log"
   "database/sql"
   "os"

   "go_webserver/api"

   _ "github.com/mattn/go-sqlite3"
 )
 

const file string = "./db/main.db"

type Users struct {
  db *sql.DB
}

func get_sql(file string) (string){
  b, err := os.ReadFile("sql/" + file + ".sql")
  if err != nil {
    return " "
  }
  str := string(b)
  return str
}


func MakeConn() (*Users, error){
  db, err :=sql.Open("sqlite3", file)
  if err != nil {
    return nil, err 
  }
  create := get_sql("create_users")
  if _, err := db.Exec(create); err != nil {
    return nil, err
  }
  return &Users {
    db:db, 
  }, nil
}


func(c *Users) Insert(user api.UserData) (int, error) {
  insert := get_sql("insert_user")
  res, err := c.db.Exec(insert, user.FName, user.LName, user.Phone, user.Email, user.Time)
  if err != nil {
    return 0, err
  }
  var id int64
  if id, err = res.LastInsertId(); err != nil {
    return 0, err
  }
  log.Printf("Added %v as %d", user, id)
  return int(id), nil
}


func (c *Users) Retrieve(name string) (api.UserData, error){  
  log.Printf("Getting %s", name)
  row := c.db.QueryRow("select * from users where FName = ?", name)
  user := api.UserData{}
  var err error
  if err = row.Scan(&user.ID, &user.FName, &user.LName, &user.Phone, &user.Email, &user.Time); err == sql.ErrNoRows {
    log.Printf("Couldn't Find User")
    return api.UserData{}, err
  }
  return user, err
}


