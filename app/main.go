package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"goFiber/internal/routers"
	"log"
)

func main() {
	engine := html.New("./web/templates", ".html") // Load templates

	app := fiber.New(fiber.Config{
		Views: engine, // Load templates
	})

	// Serve static files
	app.Static("/assets", "./web/assets")

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Routes
	routers.HTMLRendering(app)

	log.Fatal(app.Listen(":3003"))
}
