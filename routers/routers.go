package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Init() *fiber.App {
	engine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"title_website":  "GOLANG WEBSITE",
			"body_website":   "GOLANG BODY",
			"footer_website": "GOLANG FOOTER",
		})
	})

	return app
}
