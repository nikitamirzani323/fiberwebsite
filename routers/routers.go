package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/nikitamirzani323/fiberwebsite/controller"
)

func Init() *fiber.App {
	engine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/assets", "./public/assets", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})
	app.Get("/", controller.Home)

	return app
}
