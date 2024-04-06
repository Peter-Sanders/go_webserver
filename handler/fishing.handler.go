package handler 

 import (
	"github.com/Peter-Sanders/go_webserver/view/main_view"
	"github.com/labstack/echo/v4"

)

func NewFishingHandler() *FishingHandler {

	return &FishingHandler{
	}
}

type FishingHandler struct {
}


func (fh *FishingHandler) fishingHandler(c echo.Context) error {
	fishingView := main_view.Fishing()
	isError = false

	return renderView(c, main_view.FishingIndex(
		"| Fishing",
		isError,
		fishingView,
	))
}
