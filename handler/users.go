package handler

import (
  "github.com/labstack/echo/v4"

  "github.com/Peter-Sanders/go_webserver/view/user"
  "github.com/Peter-Sanders/go_webserver/model"
)


type UserHandler struct{

}

func(h UserHandler) HandleUserShow(c echo.Context) error {
  u := model.User{
    Email: "123@456.com",
  }
  return render(c, user.Show(u))
}
