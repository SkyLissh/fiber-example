package routes

import (
	"skylissh/fiber-example/crud"
	"skylissh/fiber-example/models"
	"skylissh/fiber-example/schemas"

	"github.com/gofiber/fiber/v2"
)

func AlbumRoute(router fiber.Router) {
	router.Get("/", getAlbums)
	router.Get("/:id", getAlbumByID)

	router.Post("/", postAlbums)
	router.Patch("/:id", patchAlbum)

	router.Delete("/:id", deleteAlbum)
}

// Get Albums 		godoc
// @Summary      Show all albums
// @Description  Get all albums from the store
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Success      200  {array}   []models.Album
// @Failure      500  {object}  string
// @Router       /v1/albums [get]
func getAlbums(c *fiber.Ctx) error {
	albums := new([]models.Album)

	if err := crud.Album.Find(albums); err != nil {
		return c.Status(500).SendString("Server error")
	}

	return c.JSON(albums)
}

// Get Album By ID		godoc
// @Summary      Get an album by ID
// @Description  Get an album by ID
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Album ID"
// @Success      200  {object}  models.Album
// @Failure      404  {object}  string
// @Router       /v1/albums/{id} [get]
func getAlbumByID(c *fiber.Ctx) error {
	id := c.Params("id")
	album := new(models.Album)

	if err := crud.Album.FindByID(id, album); err != nil {
		return c.Status(404).SendString("Album not found")
	}

	return c.Status(200).JSON(album)
}

// Post Albums			godoc
// @Summary      Create an album
// @Description  Create an album in the store
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Param        body  body      schemas.Album  true  "Album object"
// @Success      201   {object}  models.Album
// @Failure      400   {object}  string
// @Failure      500   {object}  string
// @Router       /v1/albums [post]
func postAlbums(c *fiber.Ctx) error {
	album := new(models.Album)
	updated := new(schemas.Album)

	if err := c.BodyParser(updated); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	if err := crud.Album.Insert(album, updated); err != nil {
		return c.Status(500).SendString("Server error")
	}

	return c.Status(200).JSON(album)
}

// Patch Albums			godoc
// @Summary      Update an album
// @Description  Update an album in the store
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Param        id    path      string         true  "Album ID"
// @Param        body  body      schemas.Album  true  "Album object"
// @Success      200   {object}  models.Album
// @Failure      400   {object}  string
// @Failure      404   {object}  string
// @Router       /v1/albums/{id} [patch]
func patchAlbum(c *fiber.Ctx) error {
	id := c.Params("id")

	data := new(schemas.Album)
	album := new(models.Album)

	if err := crud.Album.FindByID(id, album); err != nil {
		return c.Status(404).SendString("Album not found")
	}

	if err := c.BodyParser(data); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	return c.Status(200).JSON(album)
}

// Delete Albums			godoc
// @Summary      Delete an album
// @Description  Delete an album from the store
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Album ID"
// @Success      204  {object}  string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/albums/{id} [delete]
func deleteAlbum(c *fiber.Ctx) error {
	id := c.Params("id")

	album := new(models.Album)
	
	if err := crud.Album.FindByID(id, album); err != nil {
		return c.Status(404).SendString("Album not found")
	}

	if err := crud.Album.Remove(id, album); err != nil {
		return c.Status(500).SendString("Server error")
	}
	
	return c.Status(204).SendString("Album deleted")
}
