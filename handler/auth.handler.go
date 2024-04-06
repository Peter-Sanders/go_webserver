package handler 

import (
	// "fmt"
	// "errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Peter-Sanders/go_webserver/service"
  "github.com/Peter-Sanders/go_webserver/view/admin_view"
  "crypto/sha256"
  "encoding/hex"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
  auth_sessions_key string = "authenticate-sessions"
	auth_key          string = "authenticated"
	user_id_key       string = "user_id"
	username_key      string = "username"
	tzone_key         string = "time_zone"
)

type AdminService interface {
	Login(username string) (service.User, error)
}

func NewAdminHandler(as AdminService) *AdminHandler {

	return &AdminHandler{
		AdminServices: as,
	}
}


type AdminHandler struct {
	AdminServices AdminService
}


func (ah *AdminHandler) loginHandler(c echo.Context) error {
	loginView := admin_view.Login()
	isError = false

	if c.Request().Method == "POST" {
		// obtaining the time zone from the POST request of the login form
		tzone := ""
		if len(c.Request().Header["X-Timezone"]) != 0 {
			tzone = c.Request().Header["X-Timezone"][0]
		}

		// Authentication goes here
		user, err := ah.AdminServices.Login(c.FormValue("username"))
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {

				return c.Redirect(http.StatusSeeOther, "/login")
			}

			return echo.NewHTTPError(
				echo.ErrInternalServerError.Code,
				fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
		}

    inpwd :=c.FormValue("password")
    h := sha256.New()
    h.Write([]byte(inpwd))
    hpwd := hex.EncodeToString(h.Sum(nil))
    if hpwd != user.Password {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		sess, _ := session.Get(auth_sessions_key, c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600, // in seconds
			HttpOnly: true,
		}

		sess.Values = map[interface{}]interface{}{
			auth_key:     true,
			user_id_key:  user.ID,
			username_key: user.Username,
			tzone_key:    tzone,
		}
		sess.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusSeeOther, "/admin/create")
	}

	return renderView(c, admin_view.LoginIndex(
		"| Login",
		isError,
		loginView,
	))
}


func (ah *AdminHandler) adminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get(auth_sessions_key, c)
		if auth, ok := sess.Values[auth_key].(bool); !ok || !auth {
			fromProtected = false

			return echo.NewHTTPError(echo.ErrUnauthorized.Code, "Please provide valid credentials")
		}

		if userId, ok := sess.Values[user_id_key].(int); ok && userId != 0 {
			c.Set(user_id_key, userId) // set the user_id in the context
		}

		if username, ok := sess.Values[username_key].(string); ok && len(username) != 0 {
			c.Set(username_key, username) // set the username in the context
		}

		if tzone, ok := sess.Values[tzone_key].(string); ok && len(tzone) != 0 {
			c.Set(tzone_key, tzone) // set the client's time zone in the context
		}

		fromProtected = true

		return next(c)
	}
}
