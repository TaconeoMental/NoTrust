package middleware

import (
	"notrust-server/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckAuthMiddleware(c *fiber.Ctx) error {
	ntid := c.Cookies("NTID")
	if ntid == "" {
		return fiber.ErrUnauthorized
	}

	err := utils.ValidateJWT(ntid)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

