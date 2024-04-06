 package handler

 import (
	"github.com/Peter-Sanders/go_webserver/view/main_view"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)



func NewHomeHandler() *HomeHandler {

	return &HomeHandler{
	}
}

type HomeHandler struct {
}


func (hh *HomeHandler) homeHandler(c echo.Context) error {
	homeView := main_view.Home()
	isError = false

	return renderView(c, main_view.HomeIndex(
		"| Home",
		isError,
		homeView,
	))
}



func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

