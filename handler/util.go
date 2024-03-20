package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, componenet templ.Component) error {
  return componenet.Render(c.Request().Context(), c.Response())
}