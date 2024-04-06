package handler 

 import (
	"github.com/Peter-Sanders/go_webserver/view/main_view"
	"github.com/labstack/echo/v4"

)

func NewMiscHandler() *MiscHandler {

	return &MiscHandler{
	}
}

type MiscHandler struct {
}


func (ch *MiscHandler) miscHandler(c echo.Context) error {
	miscView := main_view.Misc()
	isError = false

	return renderView(c, main_view.MiscIndex(
		"| Coding",
		isError,
		miscView,
	))
}
