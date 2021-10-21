package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) http.Handler {
	e := echo.New()
	e.Use(middleware.CORS())

	h := &handler {
		DB: db,
	}

	e.GET("/health", h.health)
	e.GET("/todo", h.createTodo)
	e.GET("/todo", h.getTodoLists)
	e.DELETE("/todo/:id", h.deleteTodo)

	return e
}

func (h *handler) health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}