package app

import (
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()
	app.Listen(":3000")
}
