package handlers

import (
	"notrust-server/internal/middleware"
	"notrust-server/internal/utils"
	"notrust-server/internal/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func ServeLoginPage(c *fiber.Ctx) error {
	var methods = []models.AuthMethod{
		models.AuthMethod{"local"},
	};
    return c.Render("logintest", fiber.Map{
        "AuthMethods": methods,
    })
	//return c.SendFile("templates/public/login.html")
}

func HandleLoginRedirect(c *fiber.Ctx) error {
	ntredirect := c.Cookies("NTREDIRECT")
	if ntredirect == "" {
		return fiber.ErrBadRequest
	}

	tokenString, err := utils.GenerateJWT()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Cookie(&fiber.Cookie{
		Name:     "NTID",
		Value:    tokenString,
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	middleware.AuditLog(c, "redirect_after_login", "Redireccionado a "+ntredirect)

	return c.Redirect(ntredirect)
}

