package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/mbaghershahsavani-gif/kingdns/apps/controller/internal/models"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {

	app.Get("/api/nodes", func(c *fiber.Ctx) error {
		var nodes []models.Node
		db.Find(&nodes)
		return c.JSON(nodes)
	})

	app.Post("/api/nodes/register", func(c *fiber.Ctx) error {
		node := new(models.Node)

		if err := c.BodyParser(node); err != nil {
			return err
		}

		if node.Status == "" {
			node.Status = "online"
		}

		db.Create(node)

		return c.JSON(node)
	})

	app.Get("/api/services", func(c *fiber.Ctx) error {
		var services []models.Service
		db.Find(&services)
		return c.JSON(services)
	})

	app.Post("/api/services", func(c *fiber.Ctx) error {
		service := new(models.Service)

		if err := c.BodyParser(service); err != nil {
			return err
		}

		db.Create(service)

		return c.JSON(service)
	})
}
