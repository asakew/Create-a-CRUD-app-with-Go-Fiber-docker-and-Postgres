package routers

import "github.com/gofiber/fiber/v2"

func HTMLRendering(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":       "Go Fiber Template Example",
			"Description": "An example template",
			"Greeting":    "Hello, world!",
		})
	})

	// Facts routes
	app.Get("/createFact", func(c *fiber.Ctx) error {
		return c.Render("createFact", nil)
	})
	app.Get("/editFact", func(c *fiber.Ctx) error {
		return c.Render("editFact", nil)
	})

}
