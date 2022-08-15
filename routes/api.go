package routes

import (
	routes "skylissh/fiber-example/routes/endpoints"

	"github.com/gofiber/fiber/v2"
)

func API(router fiber.Router) {
	routes.HelloRoute(router.Group("/").Name("hello"))
	routes.AlbumRoute(router.Group("/albums").Name("album"))
	routes.UrlRoute(router.Group("/urls").Name("url"))
}
