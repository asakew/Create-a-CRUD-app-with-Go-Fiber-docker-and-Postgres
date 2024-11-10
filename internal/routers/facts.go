package routers

import (
	"github.com/gofiber/fiber/v2"
	"goFiber/internal/handlers"
)

func FactsRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView)

	app.Post("/fact", handlers.CreateFact)
	app.Get("/fact/:id/edit", handlers.EditFact)
	app.Put("/fact/:id", handlers.UpdateFact)
	app.Delete("/fact/:id", handlers.DeleteFact)

	// Add new route to show single Fact, given `:id`
	app.Get("/fact/:id", handlers.ShowFact)
}
