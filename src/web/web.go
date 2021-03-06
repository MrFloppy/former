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
	e.GET("/", index)
	e.GET("/testmail", testmail)
	e.POST("/x/*", sendMail)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "")
}

func testmail(c echo.Context) error {
	return c.Render(http.StatusOK, "mailform", "")
}

func sendMail(c echo.Context) error {

	post_data := map[string]interface{}{
		"sender":  c.FormValue("sender"),
		"email":   c.FormValue("mail"),
		"subject": c.FormValue("subject"),
		"message": c.FormValue("message"),
	}

	return c.Render(http.StatusOK, "mailsent", post_data)
}
