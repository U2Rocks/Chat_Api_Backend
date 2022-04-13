package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Tester(c *fiber.Ctx) error {
	return c.SendString("Hello Go!")
}
