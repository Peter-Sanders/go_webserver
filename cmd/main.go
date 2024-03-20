package main

import (
	// "fmt"
  "context"

	handler "github.com/Peter-Sanders/go_webserver/handler"
	"github.com/labstack/echo/v4"
)

func main()  {
  app := echo.New()

  userHandler := handler.UserHandler{}
  app.Use(withUser)
  app.GET("/user", userHandler.HandleUserShow)
  // app.GET("/")

  app.Start(":8081")
}


func withUser(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    c.Set("user", "123@456.com")
    ctx := context.WithValue(c.Request().Context(), "user", "123@456.com")
    c.SetRequest(c.Request().WithContext(ctx))
    return next(c)
  }
}

