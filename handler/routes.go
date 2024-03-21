package handler

import (

	"github.com/labstack/echo/v4"
)

var (
	fromProtected bool = false
	isError       bool = false
)

func SetupRoutes(e *echo.Echo, ah *AuthHandler, nr *NotReadyHandler) {
	e.GET("/", ah.homeHandler)
	e.GET("/login", ah.loginHandler)
	e.POST("/login", ah.loginHandler)
	e.GET("/register", ah.registerHandler)
	e.POST("/register", ah.registerHandler)

	protectedGroup := e.Group("/notready", ah.authMiddleware)
  protectedGroup.GET("/comebacklater", nr.notreadyHandler)
	/* ↓ Protected Routes ↓ */
	// protectedGroup.GET("/list", th.todoListHandler)
	// protectedGroup.GET("/create", th.createTodoHandler)
	// protectedGroup.POST("/create", th.createTodoHandler)
	// protectedGroup.GET("/edit/:id", th.updateTodoHandler)
	// protectedGroup.POST("/edit/:id", th.updateTodoHandler)
	// protectedGroup.DELETE("/delete/:id", th.deleteTodoHandler)
	// protectedGroup.POST("/logout", th.logoutHandler)
}
