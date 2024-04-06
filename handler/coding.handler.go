package handler 

 import (
	"github.com/Peter-Sanders/go_webserver/view/main_view"
	"github.com/labstack/echo/v4"

)

func NewCodingHandler() *CodingHandler {

	return &CodingHandler{
	}
}

type CodingHandler struct {
}


func (ch *CodingHandler) codingHandler(c echo.Context) error {
	codingView := main_view.Coding()
	isError = false

	return renderView(c, main_view.CodingIndex(
		"| Coding",
		isError,
		codingView,
	))
}
