package route

import (
	"loly-eiei/render"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoute(app *fiber.App, render *render.Render) {
	app.Get("/", render.HomePage)
	app.Get("/about", render.AboutPage)
}
