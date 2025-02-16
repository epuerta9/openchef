package handlers

import (
	"github.com/epuerta9/openchef/internal/database"
	"github.com/epuerta9/openchef/web/templates/pages"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	db *database.DB
}

func New(db *database.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) HandleHome(c echo.Context) error {
	component := pages.Home()
	return component.Render(c.Request().Context(), c.Response().Writer)
}
