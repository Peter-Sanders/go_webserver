package handler

import (

	"github.com/labstack/echo/v4"
)

var (
	fromProtected bool = false
	isError       bool = false
)

func SetupRoutes(e *echo.Echo, hh *HomeHandler, fh *FishingHandler, ch *CodingHandler, mh *MiscHandler, ah *AdminHandler, conh * ContentHandler) {
	e.GET("/", hh.homeHandler)
  e.GET("/home", hh.homeHandler)
  e.GET("/fishing",fh.fishingHandler)
  e.GET("/coding",ch.codingHandler)
  e.GET("/misc", mh.miscHandler)
  e.GET("login", ah.loginHandler)
  e.POST("login", ah.loginHandler)

  protectedGroup := e.Group("/admin", ah.adminMiddleware)
  protectedGroup.GET("/create", conh.createContentHandler)
}
