package render

import "github.com/gofiber/fiber/v2"

func (r *Render) AboutPage(c *fiber.Ctx) error {
	conf := r.conf
	return c.Render("AboutPage", fiber.Map{
		"Title":       "About Page",
		"FileVersion": conf.Server.FileVersion,
	}, "layouts/main")
}
