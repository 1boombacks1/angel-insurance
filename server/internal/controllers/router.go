package controllers

import (
	"net/http"

	"github.com/1boombacks1/botInsurance/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(app *fiber.App, services *services.Services) {
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	api := app.Group("/api")
	{
		newClientRoutes(api.Group("/client"), services.CarInsuranceApplication)
	}
}
