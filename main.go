 package main

 import (
   "text/template"
   "io"
   "net/http"
   "log"
   "time"

   "go_webserver/server"
   "go_webserver/api"

   "github.com/labstack/echo/v4"
   _ "github.com/mattn/go-sqlite3"
 )

type Template struct {
   templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}


 func main() {
    // var err error
    e := echo.New()

    var user *server.Users 
    var err error 
    if user, err = server.MakeConn(); err != nil {
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
    e.GET("/", func(c echo.Context) error {
      return c.Render(http.StatusOK, "index", nil)
    })
    
    e.GET("/new_user", func(c echo.Context) error {
      log.Printf("Start Making New User")
      return c.Render(http.StatusOK, "new_user", nil)  
    })

    e.GET("/name_search", func(c echo.Context) error {
      return c.Render(http.StatusOK, "name_search", nil)
    })

    e.GET("/list_users", func(c echo.Context) error{
      users, err := user.List()
      if err != nil{
        log.Print("we have no users?")
      }
      log.Print(users)
      return c.Render(http.StatusOK, "list", users)
    })

    e.POST("/create_user", func(c echo.Context) error {
      fname := c.FormValue("FName")
      lname := c.FormValue("LName")
      phone := c.FormValue("Phone")
      email := c.FormValue("Email")
      time := time.Now().Format("2022-03-05 20:20:20")

      var data api.UserData
      data.FName = fname
      data.LName = lname
      data.Phone = phone
      data.Email = email 
      data.Time = time 
      data.ID = 0
      log.Printf("%v", data)

      id, err := user.Insert(data)
      if err != nil {
        log.Fatal(err)
      }
      log.Printf("Created new %v with id %d", data, id)
      return c.Render(http.StatusOK, "new_user", nil) 
    })

    e.PUT("/user-data", func(c echo.Context) error {

      name := c.FormValue("FName")
      if name =="" {
        log.Printf("No user passed")
        return c.Render(http.StatusOK, "no_user_passed", nil)
      }
      user_data, err := user.Retrieve(name)
      if err != nil {
        // log.Fatal(err)
        return c.Render(http.StatusOK, "user_not_found", nil)
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
