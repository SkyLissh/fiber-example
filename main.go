package main

import (
	"skylissh/fiber-example/core"
	"skylissh/fiber-example/docs"
	"skylissh/fiber-example/routes"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @version     1.0
// @description This is a sample API for using Fiber
func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Swagger Endpoint
	docs.SwaggerInfo.Title = core.Settings.ProjectName

	app.Get("/docs/*", swagger.HandlerDefault)

	// Routes
	api := app.Group(core.Settings.ApiVersion).Name("api")
	routes.API(api)

	app.Listen(":8000")
}
