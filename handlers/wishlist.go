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
	name := c.QueryParam("wish")
	description := c.QueryParam("description")
	link := c.QueryParam("link")

	wish := store.Wish{}

	err := h.DB.QueryRow("INSERT INTO wishlist (priority, wish, description, link) VALUES ($1, $2, $3, $4) RETURNING id, priority, wish, description, link", priority, name, description, link).
		Scan(&wish.ID, &wish.Priority, &wish.Wish, &wish.Description, &wish.Link)
	if err != nil {
		log.Printf("database query error: %s", err)
		return err
	}

	return c.JSON(http.StatusOK, wish)
}

func (h *Handler) DeleteWish(c echo.Context) error {
	id := c.QueryParam("id")

	_, err := h.DB.Exec("DELETE FROM wishlist WHERE id = $1", id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
