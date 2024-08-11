package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	// home page
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", nil)
	})

	// signup and login routes
	e.GET("/signup", func(c echo.Context) error {
		return c.Render(200, "signup", nil)
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(200, "login", nil)
	})

	// auth protected routes
	e.GET("/profile", func(c echo.Context) error {
		return c.Render(200, "profile.html", nil)
	})

	// 404 page -> all not found routes
	e.GET("/*", func(c echo.Context) error {
		return c.Render(404, "404", nil)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
