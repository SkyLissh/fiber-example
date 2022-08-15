package routes

import "github.com/gofiber/fiber/v2"

func HelloRoute(router fiber.Router) {
	router.Get("/", sayHello)
}

// Say Hello 		godoc
// @Summary     Say hello
// @Description Say hello to the world
// @Tags        Hello
// @Accept      json
// @Produce     json
// @Success     200 {object} string
// @Router      /v1/ [get]
func sayHello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, world!",
	})
}
