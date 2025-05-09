package handlers

import (
	"notrust-server/internal/middleware"
	"notrust-server/internal/models"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var body models.LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	if body.Password != "123456" {
		middleware.AuditLog(c, "login_failure", "Login fallido para "+body.Email)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Contraseña incorrecta."})
	}

	middleware.AuditLog(c, "login_success", "Login exitoso para "+body.Email)
	return c.JSON(fiber.Map{"success": true})
}

func OtpHandler(c *fiber.Ctx) error {
	var body models.OTPRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	if body.OTP != "654321" {
		middleware.AuditLog(c, "otp_failure", "OTP inválido ingresado")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Código OTP incorrecto."})
	}

	middleware.AuditLog(c, "otp_success", "OTP correcto")
	return c.JSON(fiber.Map{"success": true})
}

func SendOtpHandler(c *fiber.Ctx) error {
	var body models.SendOTPRequest
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	middleware.AuditLog(c, "otp_sent", "OTP enviado a "+body.Email)
	return c.JSON(fiber.Map{"success": true})
}

