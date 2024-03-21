package handler

import (
  "github.com/Peter-Sanders/go_webserver/view/notready"

  "github.com/labstack/echo/v4"
)

type NotReadyHandler struct {
  test string
}

func NewNotReadyHandler(test string) *NotReadyHandler {

  return &NotReadyHandler{
    test: test,
  }
}

func (nr *NotReadyHandler) notreadyHandler(c echo.Context) error {
	homeView := notready.Temp(fromProtected)
	isError = false

	return renderView(c, notready.TempIndex(
		"| Not Ready",
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		homeView,
	))
}

// func renderView(c echo.Context, cmp templ.Component) error {
// 	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
//
// 	return cmp.Render(c.Request().Context(), c.Response().Writer)
// }
