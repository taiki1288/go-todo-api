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