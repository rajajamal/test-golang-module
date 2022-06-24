package main

import (
	"fiber/configs"
	"fiber/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

func init() {
	godotenv.Load()
	configs.Load()
	configs.Db.AutoMigrate(
		&models.User{},
	)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("مرحبا بالعالم")
	})

	app.Listen(":3000")
}
