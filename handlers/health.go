package handlers

import (
	"github.com/ktiutiun/plan.git/store"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *Handler) GetHealthHabits(c echo.Context) error {
	healthHabits, err := store.GetHealthHabits(h.DB)
	if err != nil {
		log.Println("Помилка отримання значень шкал:", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.Render(http.StatusOK, "health.html", healthHabits)
}

func (h *Handler) AddHealthHabits(c echo.Context) error {
	scale := c.QueryParam("scale")
	habit := c.QueryParam("habit")

	_, err := h.DB.Exec("INSERT INTO health_habits (habit_name, scale) VALUES ($1, $2) ON CONFLICT (habit_name) DO UPDATE SET scale = EXCLUDED.scale", habit, scale)
	if err != nil {
		log.Printf("database query error: %s", err)
		return err
	}

	return c.NoContent(http.StatusOK)
}
