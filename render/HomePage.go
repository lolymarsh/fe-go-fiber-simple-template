package render

import "github.com/gofiber/fiber/v2"

func (r *Render) HomePage(c *fiber.Ctx) error {
	conf := r.conf
	return c.Render("HomePage", fiber.Map{
		"Title":       "Home Page",
		"FileVersion": conf.Server.FileVersion,
	}, "layouts/main")
}
