package main

import (
	// "fmt"
  // "context"
	"github.com/Peter-Sanders/go_webserver/handler"
  // "github.com/Peter-Sanders/go_webserver/db"
  // "github.com/Peter-Sanders/go_webserver/service"

	// "github.com/gorilla/sessions"
	// "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

const (
  SECRET_KEY string = "secret"
  DB_NAME string = "./db/main.db"
)
 
func main()  {
  app := echo.New()
  app.Static("/static", "static")

  app.HTTPErrorHandler = handler.CustomHTTPErrorHandler
  // app.Use(middleware.Logger())
  // app.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))

  // store, err := db.NewStore(DB_NAME)
  // if err != nil {
  //   app.Logger.Fatalf("failed to create store: %s", err)
  // }

  // us := service.NewUserServices(service.User{}, store)
	hh := handler.NewHomeHandler()


	// Setting Routes
	handler.SetupRoutes(app, hh)
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
