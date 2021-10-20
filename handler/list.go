package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type List struct {
	ID        uint      `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"` 
}

func (h *handler) getTodoLists(c echo.Context) error {
	var lists []List
	err := h.DB.Find(&lists).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lists)
}

func (h *handler) createTodo(c echo.Context) error {
	var list List
	err := c.Bind(&list)
	if err != nil {
		return err
	}

	err = h.DB.Create(&list).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, list)
}