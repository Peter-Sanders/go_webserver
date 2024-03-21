package handler

import (
	"fmt"
	"net/http"

	"github.com/Peter-Sanders/go_webserver/view/error_page"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	var errorPage func(fp bool) templ.Component

	switch code {
	case 401:
		errorPage = error_page.Error401
	case 404:
		errorPage = error_page.Error404
	case 500:
		errorPage = error_page.Error500
	}

	isError = true

	renderView(c, error_page.ErrorIndex(
		fmt.Sprintf("| Error (%d)", code),
		"",
		fromProtected,
		isError,
		errorPage(fromProtected),
	))
}
