package handlers

import (
	"github.com/ktiutiun/plan.git/store"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *Handler) GetWishlist(c echo.Context) error {
	wishes, err := store.GetWishes(h.DB)
	if err != nil {
		log.Println("Помилка отримання бажань:", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Render(http.StatusOK, "wishlist.html", wishes)
}

func (h *Handler) AddWishlist(c echo.Context) error {
	priority := c.QueryParam("priority")
	wish := c.QueryParam("wish")
	description := c.QueryParam("description")
	link := c.QueryParam("link")

	_, err := h.DB.Exec("INSERT INTO wishlist (priority, wish, description, link) VALUES ($1, $2, $3, $4)", priority, wish, description, link)
	if err != nil {
		log.Printf("database query error: %s", err)
		return err
	}

	return c.NoContent(http.StatusOK)
}
