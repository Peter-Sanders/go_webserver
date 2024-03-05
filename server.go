 package main

 import (
   "text/template"
   "io"
   "net/http"
   // "sync"
   // "time"
   "log"
   "database/sql"
   // "handler/handler"

   "github.com/labstack/echo/v4"
   _ "github.com/mattn/go-sqlite3"
 )

type Template struct {
   templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

type Users struct {
  db *sql.DB
}

type UserData struct {
  FName string `json:"FName"`
  LName string `json:"LName"`
  Phone string `json:"Phone"`
  Email string `json:"Email"`
  Time string  `json:"Time"`
  ID int       `json:"ID"`
}

func HomeHandler(c echo.Context) error {
  // Please note the the second parameter "home.html" is the template name and should
  // be equal to the value stated in the {{ define }} statement in "view/home.html"
  return c.Render(http.StatusOK, "home.html", map[string]interface{}{
    "name": "HOME",
    "msg": "Hello, Boatswain!",
  })
}

const file string = "./db/main.db"

const create string = `
create table if not exists users (
  id integer not null primary key,
  FName text,
  LName text,
  Phone text,
  Email text,
  Time text);
`
const insert string = `
insert into users 
with data as (
  select 
  NULL as id, 
  "Bob" as fname, 
  "White" as lname, 
  "(420) 867-5309" as phone, 
  "bwhite@quail.com" as email, 
  "2022-03-05 00:00:00" as time 
)
select 
  d.*
from data d 
left join users u on u.email = d.email and lower(u.fname) = lower(d.fname) and lower(u.lname) = lower(d.lname)
where u.id is null;`

func MakeConn() (*Users, error){
  db, err :=sql.Open("sqlite3", file)
  if err != nil {
    return nil, err 
  }
  if _, err := db.Exec(create); err != nil {
    return nil, err
  }
  if _, err := db.Exec(insert); err != nil{
    return nil, err
  }
  return &Users {
    db:db, 
  }, nil
}

func (c *Users) Retrieve(name string) (UserData, error){  
  log.Printf("Getting %s", name)
  row := c.db.QueryRow("select * from users where FName = ?", name)
  user := UserData{}
  var err error
  if err = row.Scan(&user.ID, &user.FName, &user.LName, &user.Phone, &user.Email, &user.Time); err == sql.ErrNoRows {
    log.Printf("Couldn't Find User")
    return UserData{}, err
  }
  return user, err
}



 func main() {
    // var err error
    e := echo.New()

    var user *Users 
    var err error 
    if user, err = MakeConn(); err != nil {
      log.Fatal(err)
    }

    renderer :=&Template{
      templates: template.Must(template.ParseGlob("static/*.html")),
    }
    e.Renderer = renderer
    e.Static("/styles", "styles")


    // e.Pre(middleware.RemoveTrailingSlash())
    // e.Use(middleware.Recover())
    // e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
    //   rate.Limit(20),
    // )))
    // template.NewTemplateRenderer(e, "static/*.html")
    e.GET("/", func(c echo.Context) error {
      // res := map[string]interface{}{
      //   "Name": "Pete",
      //   "Phone": "867-5309",
      //   "Email": "testing@testing.com",
      // }
      return c.Render(http.StatusOK, "index", nil)
    })
    e.GET("/get-info", func(c echo.Context) error {
      res := map[string]interface{}{
        "Name": "Pete",
        "Phone": "867-5309",
        "Email": "testing@testing.com",
      }
      return c.Render(http.StatusOK, "name_card", res)
    })
    e.POST("/user-data", func(c echo.Context) error {
      // db, err := MakeConn()
      // if err != nil {
      //   return nil
      // }
      // db, err :=sql.Open("sqlite3", file)

      name := c.FormValue("FName")
      log.Printf("%s", name)
      user_data, err := user.Retrieve(name)
      if err != nil {
        log.Fatal(err)
      }
      res := map[string]interface{}{
        "Name": user_data.FName + " " + user_data.LName,
        "Phone": user_data.Phone,
        "Email": user_data.Email,
      }
      return c.Render(http.StatusOK, "name_card", res)


    })

   e.Logger.Fatal(e.Start(":8081"))
 }



 // func main() {
 //     http.Handle("/", http.FileServer(http.Dir("./static")))
 //
 //   // log.Fatal(http.ListenAndServe(":8081", nil))
 //   log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
 //
 // }
