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

// Render renders the given template with the provided data to the given io.Writer.
// It takes the name of the template, the data to be rendered, and the echo.Context as parameters.
// It returns an error.
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("../views/*.html")),
	}
}

type Block struct {
	Id int
}

type Blocks struct {
	Start  int
	Next   int
	More   bool
	Blocks []Block
}

type Count struct {
	Count int
}

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "jd@gmail,com"),
			newContact("Jane", "jd@gmail,com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	page := newPage()
	e.Renderer = NewTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			return c.Render(422, "form", formData)
		}

		contact := newContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)

		c.Render(200, "form", newFormData())
		return c.Render(200, "oob-contact", contact)
	})

	//e.GET("/blocks", func(c echo.Context) error {
	//	startStr := c.QueryParam("start")
	//	start, err := strconv.Atoi(startStr)
	//	if err != nil {
	//		start = 0
	//	}
	//
	//	blocks := []Block{}
	//	for i := start; i < start+10; i++ {
	//		blocks = append(blocks, Block{Id: i})
	//	}
	//
	//	template := "blocks"
	//	if start == 0 {
	//		template = "blocks-index"
	//	}
	//	return c.Render(http.StatusOK, template, Blocks{
	//		Start:  start,
	//		Next:   start + 10,
	//		More:   start+10 < 100,
	//		Blocks: blocks,
	//	})
	//})

	e.Logger.Fatal(e.Start(":42069"))
}
