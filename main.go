package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! From Daew project")
	})

	app.Get("/say-hi/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Name parameter is required")
		}
		return c.SendString("Hello, " + name)
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		if name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Name is required")
		}
		return c.SendString("Form submitted successfully with name: " + name)
	})

	app.Post("/json", func(c *fiber.Ctx) error {
		type RequestBody struct {
			Name string `json:"name"`
		}
		var body RequestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
		}
		if body.Name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Name is required")
		}
		return c.JSON(fiber.Map{
			"message": "JSON received successfully",
			"name":    body.Name,
		})
	})

	app.Listen(":8080")
}
