package main

import (
	"delta-go/pkg/common/db"
	"delta-go/pkg/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	h := db.Init()
	app := fiber.New()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	users.RegisterRoutes(app, h)

	app.Listen("8080")
}
