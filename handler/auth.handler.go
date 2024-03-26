package handler 

import (
	// "fmt"
	// "errors"
	// "fmt"
	// "net/http"
	// "strings"

	"github.com/Peter-Sanders/go_webserver/service"
	"github.com/Peter-Sanders/go_webserver/view/auth_view"
	// "golang.org/x/crypto/bcrypt"   

	"github.com/a-h/templ"
	// "github.com/gorilla/sessions"
	// "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
auth_sessions_key string = "authenticate-sessions"
	auth_key          string = "authenticated"
	user_id_key       string = "user_id"
	username_key      string = "username"
	tzone_key         string = "time_zone"
)

type AuthService interface {
	CreateUser(u service.User) error
	CheckEmail(email string) (service.User, error)
}

func NewAuthHandler(us AuthService) *AuthHandler {

	return &AuthHandler{
		UserServices: us,
	}
}

type AuthHandler struct {
	UserServices AuthService
}

func (ah *AuthHandler) homeHandlerOld(c echo.Context) error {
	homeView := auth_view.Home()
	isError = false

	return renderView(c, auth_view.HomeIndex(
		"| Home",
		isError,
		homeView,
	))
}



func renderViewOld(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

