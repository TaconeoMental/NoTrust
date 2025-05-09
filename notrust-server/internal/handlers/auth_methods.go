package handlers

import "github.com/gofiber/fiber/v2"

func GetAuthMethodsHandler(c *fiber.Ctx) error {
	//authMethods := []string{"password", "otp", "github", "google", "azure"}
	authMethods := []string{"password", "otp", "google"}

	return c.JSON(fiber.Map{
		"auth_methods": authMethods,
	})
}

