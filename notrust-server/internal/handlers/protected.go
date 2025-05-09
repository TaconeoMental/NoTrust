package handlers

import "github.com/gofiber/fiber/v2"

func ProtectedHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

