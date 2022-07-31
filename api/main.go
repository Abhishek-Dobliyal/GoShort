package main

import (
	"goshort/routes"
	"goshort/helpers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/shorten", routes.ShortenURL)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(helpers.AppPort)
}
