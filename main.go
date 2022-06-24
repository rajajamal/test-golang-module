package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rajajamal/test-golang-module/configs"
	"github.com/rajajamal/test-golang-module/models"
)

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
