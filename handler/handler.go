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

func (h *handler) health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}