package main

import (
	// "fmt"
	// "context"
	"fmt"
  "gopkg.in/ini.v1"

	"github.com/Peter-Sanders/go_webserver/db"
	"github.com/Peter-Sanders/go_webserver/handler"

	"github.com/Peter-Sanders/go_webserver/service"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


const (
  SECRET_KEY string = "secret"
)
 
func main()  {
  app := echo.New()
  app.Static("/static", "static")

  app.HTTPErrorHandler = handler.CustomHTTPErrorHandler

  inidata, err := ini.Load("env.ini")
  if err!= nil {
    fmt.Printf("Fail to read file: %v", err)
    app.Logger.Fatal("Couln't get ini data")
  }
  section := inidata.Section("database")
  DB_NAME := "db/"+section.Key("dbname").String()


  app.Use(middleware.Logger())
  app.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))

  store, err := db.NewStore(DB_NAME)
  if err != nil {
    app.Logger.Fatalf("failed to create store: %s", err)
  }
  fmt.Print(store)

  // us := service.NewUserServices(service.User{}, store)
	hh := handler.NewHomeHandler()
  fh := handler.NewFishingHandler()
  ch := handler.NewCodingHandler()
  mh := handler.NewMiscHandler()
  as := service.NewAdminServices(service.User{}, store)
  ah := handler.NewAdminHandler(as)
  conh := handler.NewContentHandler()


	// Setting Routes
	handler.SetupRoutes(app, hh, fh, ch, mh, ah, conh)
  // userHandler := handler.UserHandler{}
  // app.Use(withUser)
  // app.GET("/user", userHandler.HandleUserShow)
  // app.GET("/")

  app.Logger.Fatal(app.Start(":8081"))
}


// func withUser(next echo.HandlerFunc) echo.HandlerFunc {
//   return func(c echo.Context) error {
//     c.Set("user", "123@456.com")
//     ctx := context.WithValue(c.Request().Context(), "user", "123@456.com")
//     c.SetRequest(c.Request().WithContext(ctx))
//     return next(c)
//   }
// }
//
