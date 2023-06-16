package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"net/http"
)

type PageData struct {
	Title     string
	MenuItem1 string
	MenuItem2 string
	MenuItem3 string
}

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
		templates: template.Must(template.ParseGlob("*.html")),
	}

	e.Renderer = renderer

	templateData := &PageData{
		Title:     "Plan",
		MenuItem1: "Health",
		MenuItem2: "Wishlist",
		MenuItem3: "Budget",
	}

	e.Renderer = renderer

	// Шаблон для стартової сторінки
	e.GET("/", func(c echo.Context) error {
		templateData.Title = "Plan"
		return c.Render(http.StatusOK, "startPage.html", map[string]string{
			"page": "main",
		})
	})

	// Шаблон для сторінки "Health"
	e.GET("/health", func(c echo.Context) error {
		templateData.Title = "Health"
		return c.Render(http.StatusOK, "health.html", map[string]string{
			"page": "budget",
		})
		//return c.Render(http.StatusOK, "budget.html", templateData)
	})

	// Шаблон для сторінки "Products"
	e.GET("/wishlist", func(c echo.Context) error {
		templateData.Title = "Products"
		return c.Render(http.StatusOK, "wishlist.html", templateData)
	})

	// Шаблон для сторінки "Cart"
	e.GET("/budget", func(c echo.Context) error {
		templateData.Title = "Cart"
		return c.Render(http.StatusOK, "budget.html", templateData)
	})

	err := e.Start(":8081")
	if err != nil {
		log.Printf("start server: %s", err)
		return
	}
}
