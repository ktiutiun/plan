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
		MenuItem2: "Кошик",
		MenuItem3: "Контакти",
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "startPage.html", templateData)
	})

	e.GET("/health", func(c echo.Context) error {
		templateData.Title = templateData.MenuItem1
		return c.Render(http.StatusOK, "health.html", templateData)
	})

	e.GET("/cart", func(c echo.Context) error {
		templateData.Title = templateData.MenuItem2
		return c.Render(http.StatusOK, "startPage.html", templateData)
	})

	e.GET("/contacts", func(c echo.Context) error {
		templateData.Title = templateData.MenuItem3
		return c.Render(http.StatusOK, "startPage.html", templateData)
	})

	err := e.Start(":8080")
	if err != nil {
		return
	}
}
