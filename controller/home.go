package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/fiberwebsite/config"
)

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}
type Info struct {
	Affiliation string
	Address     string
}

func Home(c *fiber.Ctx) error {
	var layout string = "layout/template"
	var person = Person{
		Name:    "Bruce Wayne",
		Gender:  "malea",
		Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
		Info:    Info{"Wayne Enterprises", "Gotham City"},
	}
	return c.Render("home/index", fiber.Map{
		"title_website":  config.TITLE_WEBSITE,
		"image_path":     config.IMAGE_PATH,
		"record":         person,
		"footer_website": config.FOOTER_WEBSITE,
	}, layout)
}
