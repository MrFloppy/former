package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Templating
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func WebPrint() {
	fmt.Println("We are also on teh web!")
}

func RunServer() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("../public/views/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", hello)
	e.GET("/testmail", testmail)
	e.POST("/x/*", sendMail)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func testmail(c echo.Context) error {
	return c.Render(http.StatusOK, "testmail", "World")
}

func sendMail(c echo.Context) error {
	subject := c.FormValue("subject")
	message := c.FormValue("message")

	return c.String(http.StatusOK, subject+message)
}
