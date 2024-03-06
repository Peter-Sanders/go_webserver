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
  get_by_name := get_sql("get_user_by_name")
  row := c.db.QueryRow(get_by_name, name)
  user := api.UserData{}
  var err error
  if err = row.Scan(&user.ID, &user.FName, &user.LName, &user.Phone, &user.Email, &user.Time); err == sql.ErrNoRows {
    log.Printf("Couldn't Find User %s", name)
    return api.UserData{}, err
  }
  return user, err
}


func (c *Users) List() ([]api.UserData, error){
  log.Printf("Getting all users")
  list :=get_sql("get_users")
  rows, err := c.db.Query(list)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  data := []api.UserData{}
  for rows.Next() {
    i := api.UserData{}
    err = rows.Scan(&i.ID, &i.FName, &i.LName, &i.Phone, &i.Email, &i.Time)
    if err != nil {
      return nil, err
    }
    data = append(data, i)
  }
  return data, nil
}


func (c *Users) ListPaginated(limit int, offset int) ([]api.UserData, error) {
  query := get_sql("get_users_paginated")
  rows, err := c.db.Query(query, limit, offset)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  data := []api.UserData{}
  for rows.Next() {
    i := api.UserData{}
    err = rows.Scan(&i.ID, &i.FName, &i.LName, &i.Phone, &i.Email, &i.Time)
    if err != nil {
      return nil, err
    }
    data = append(data, i)
  }
  log.Print(data)
  return data, nil
}

