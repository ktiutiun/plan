package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"html/template"
	"io"
	"log"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
	db        *sql.DB
}

type Habit struct {
	Name  string
	Scale int64
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	if err := t.templates.ExecuteTemplate(w, name, data); err != nil {
		return err
	}
	return nil
}

func createDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5555 user=postgres password=postgres dbname=plan_database sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getHealthHabits(db *sql.DB) ([]Habit, error) {
	// Запит до бази даних для отримання значень шкал
	rows, err := db.Query("SELECT habit_name, scale FROM health_habits ORDER BY habit_name")
	if err != nil {
		return nil, err
	}

	habits := []Habit{}

	for rows.Next() {
		habit := Habit{}

		err := rows.Scan(&habit.Name, &habit.Scale)
		if err != nil {
			return nil, err
		}

		habits = append(habits, habit)
	}

	return habits, nil
}

func main() {
	db, err := createDBConnection()
	if err != nil {
		return
	}

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
		// Отримати значення шкал з бази даних
		healthHabits, err := getHealthHabits(db)
		if err != nil {
			log.Println("Помилка отримання значень шкал:", err)
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.Render(http.StatusOK, "health.html", healthHabits)
	})
	e.POST("/health/habits", func(c echo.Context) error {
		scale := c.QueryParam("scale")
		habit := c.QueryParam("habit")

		_, err = db.Exec("INSERT INTO health_habits (habit_name, scale) VALUES ($1, $2) ON CONFLICT (habit_name) DO UPDATE SET scale = EXCLUDED.scale", habit, scale)
		if err != nil {
			log.Printf("database query error: %s", err)
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	// Шаблон для сторінки "Wishlist"
	e.GET("/wishlist", func(c echo.Context) error {
		return c.Render(http.StatusOK, "wishlist.html", nil)
	})

	// Шаблон для сторінки "Budget"
	e.GET("/budget", func(c echo.Context) error {
		return c.Render(http.StatusOK, "budget.html", nil)
	})

	err = e.Start(":8080")
	if err != nil {
		log.Printf("start server: %s", err)
		return
	}
}
