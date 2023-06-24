package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetBudget(c echo.Context) error {
	return c.Render(http.StatusOK, "budget.html", nil)
}
