package handler 

 import (
	"github.com/Peter-Sanders/go_webserver/view/admin_view"
	"github.com/labstack/echo/v4"

)

func NewContentHandler() *ContentHandler {

	return &ContentHandler{
	}
}

type ContentHandler struct {
}


func (ch *ContentHandler) createContentHandler(c echo.Context) error {
	contentView := admin_view.CreateContent()
	isError = false

	return renderView(c, admin_view.CreateContentIndex(
		"| Create Content",
		isError,
		contentView,
	))
}
