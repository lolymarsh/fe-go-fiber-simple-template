package render

import "github.com/gofiber/fiber/v2"

func (r *Render) NotFoundPage(c *fiber.Ctx) error {
	conf := r.conf
	return c.Render("NotFoundPage", fiber.Map{
		"Title":       "Not Found",
		"FileVersion": conf.Server.FileVersion,
	}, "layouts/main")
}
