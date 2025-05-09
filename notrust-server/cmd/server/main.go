package main

import (
	"log"
	"time"
	"os"
	"notrust-server/internal/handlers"
	"notrust-server/internal/middleware"
	"notrust-server/internal/config"

	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadEnv()

    engine := html.New("./templates", ".html")
    app := fiber.New(fiber.Config{
        Views: engine,
    })

	app.Static("/static", "./static")

	// Middlewares globales
	app.Use(logger.New())
	app.Use(recover.New())

	// Public endpoints
	otp := app.Group("/api/otp")
	otp.Post("/check", limiter.New(limiter.Config{Max: 5, Expiration: 30 * time.Second}), handlers.OtpHandler)
	otp.Post("/send", limiter.New(limiter.Config{Max: 5, Expiration: 30 * time.Second}), handlers.SendOtpHandler)

	notrust := app.Group("/notrust")
	notrust.Get("/login", handlers.ServeLoginPage)
	notrust.Post("/login", handlers.HandleLoginRedirect)

	// Protected
	notrust.Get("/check", middleware.CheckAuthMiddleware, handlers.ProtectedHandler)


	log.Println("Servidor NoTrust corriendo en http://localhost:8080")
	pwd, _ := os.Getwd()
	log.Println(pwd)
	log.Fatal(app.Listen(":8081"))
}

