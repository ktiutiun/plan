package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	if err := t.templates.ExecuteTemplate(w, name, data); err != nil {
		return err
	}
	return nil
}

func main() {
	e := echo.New()

	log.Println("http://localhost:8080/")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = renderer

	// Шаблон для стартової сторінки
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "startPage.html", nil)
	})

	// Шаблон для сторінки "Health"
	e.GET("/health", func(c echo.Context) error {
		return c.Render(http.StatusOK, "health.html", nil)
	})

	// Шаблон для сторінки "Wishlist"
	e.GET("/wishlist", func(c echo.Context) error {
		return c.Render(http.StatusOK, "wishlist.html", nil)
	})

	// Шаблон для сторінки "Budget"
	e.GET("/budget", func(c echo.Context) error {
		return c.Render(http.StatusOK, "budget.html", nil)
	})

	err := e.Start(":8080")
	if err != nil {
		log.Printf("start server: %s", err)
		return
	}
}
