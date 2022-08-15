package routes

import (
	"log"
	"skylissh/fiber-example/crud"
	"skylissh/fiber-example/models"
	"skylissh/fiber-example/schemas"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UrlRoute(router fiber.Router) {
	router.Get("/", getUrls)
	router.Get("/:url_key", getUrlByKey)

	router.Post("/", postUrls)
	router.Patch("/:url_key", patchUrl)

	router.Delete("/:url_key", deleteUrl)
} 

// Get Urls 		godoc
// @Summary      Show all urls
// @Description  Get all urls from the store
// @Tags         Urls
// @Accept       json
// @Produce      json
// @Success      200  {array}   []models.Url
// @Failure      500  {object}  string
// @Router       /v1/urls [get]
func getUrls(c *fiber.Ctx) error {
	urls := new([]models.Url)

	data := crud.Url.Find(urls)
	if err := <- data; err != nil {
		return c.Status(500).SendString("Server error")
	}

	return c.JSON(urls)
}

// Get Url By Key		godoc
// @Summary      Get an url by key
// @Description  Get an url by key
// @Tags         Urls
// @Accept       json
// @Produce      json
// @Param        url_key   path      string  true  "Url key"
// @Success      200  {object}  models.Url
// @Failure      404  {object}  string
// @Router       /v1/urls/{url_key} [get]
func getUrlByKey(c *fiber.Ctx) error {
	urlKey := c.Params("url_key")
	url := new(models.Url)

	if err := crud.Url.FindByKey(urlKey, url); err != nil {
		return c.Status(404).SendString("Url not found")
	}

	return c.Status(200).JSON(url)
}

// Post Urls			godoc
// @Summary      Create a new url
// @Description  Create a new url
// @Tags         Urls
// @Accept       json
// @Produce      json
// @Param        url body schemas.Url true "Url"
// @Success      201   {object}  models.Url
// @Failure			 400				{object}  string
// @Failure      500  {object}  string
// @Router       /v1/urls [post]
func postUrls(c *fiber.Ctx) error {
	url := new(models.Url)
	body := new(schemas.Url)

	validate := validator.New()

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	if err := validate.Struct(body); err != nil {
		log.Println(err)
		return c.Status(400).SendString("Invalid body")
	}

	if err := crud.Url.FindByTargetUrl(body.TargetUrl, url); err == nil {
		return c.Status(200).JSON(url)
	}

	if err := crud.Url.Insert(url, body); err != nil {
		return c.Status(500).SendString("Server error")
	}

	return c.Status(201).JSON(url)
}

// Patch Url			godoc
// @Summary      Update an url
// @Description  Update an url in the store
// @Tags         Urls
// @Accept       json
// @Produce      json
// @Param        url_key path string true "Url key"
// @Param        url body schemas.Url true "Url"
// @Success      200  {object}  models.Url
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/urls/{url_key} [patch]
func patchUrl(c *fiber.Ctx) error {
	urlKey := c.Params("url_key")

	url := new(models.Url)
	body := new(schemas.Url)

	validate := validator.New()

	if err := crud.Url.FindByKey(urlKey, url); err != nil {
		return c.Status(404).SendString("Url not found")
	}

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	if err := validate.Struct(body); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	if err := crud.Url.Update(url.ID.String(), url, body); err != nil {
		return c.Status(500).SendString("Server error")
	}

	return c.Status(200).JSON(url)
}

// Delete Url			godoc
// @Summary      Delete an url
// @Description  Delete an url from the store
// @Tags         Urls
// @Accept       json
// @Produce      json
// @Param        url_key path string true "Url key"
// @Success      204  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/urls/{url_key} [delete]
func deleteUrl(c *fiber.Ctx) error {
	urlKey := c.Params("url_key")
	url := new(models.Url)

	if err := crud.Url.FindByKey(urlKey, url); err != nil {
		return c.Status(404).SendString("Url not found")
	}

	if err := crud.Url.Remove(url.ID.String(), url); err != nil {
		return c.Status(500).SendString("Server error")
	}

	return c.Status(204).SendString("Url deleted")
}
