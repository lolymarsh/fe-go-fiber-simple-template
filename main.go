package main

import (
	"loly-eiei/config"
	"loly-eiei/render"
	"loly-eiei/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Load config
	conf := config.LoadConfigFile()

	// Initialize the template engine
	engine := html.New("./views", ".html")
	// Enable reload for development (optional)

	engine.Reload(conf.Server.DevMode)

	// Create a new Fiber instance with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	// Initialize render and routes
	render := render.NewRender(conf)
	route.SetUpRoute(app, render)

	app.Use(func(c *fiber.Ctx) error {
		return render.NotFoundPage(c)
	})

	// Start server
	app.Listen(":" + conf.Server.PortServer)
}
