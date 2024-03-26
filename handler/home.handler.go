 package handler

 import (
	// "fmt"
	// "errors"
	// "fmt"
	// "net/http"
	// "strings"

	// "github.com/Peter-Sanders/go_webserver/service"
	"github.com/Peter-Sanders/go_webserver/view/auth_view"
	// "golang.org/x/crypto/bcrypt"   

	"github.com/a-h/templ"
	// "github.com/gorilla/sessions"
	// "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

)



func NewHomeHandler() *HomeHandler {

	return &HomeHandler{
	}
}

type HomeHandler struct {
}


func (hh *HomeHandler) homeHandler(c echo.Context) error {
	homeView := auth_view.Home()
	isError = false

	return renderView(c, auth_view.HomeIndex(
		"| Home",
		isError,
		homeView,
	))
}



func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

