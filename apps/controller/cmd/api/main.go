package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mbaghershahsavani-gif/kingdns/apps/controller/internal/database"
        "github.com/mbaghershahsavani-gif/kingdns/apps/controller/internal/handlers"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "kingdns-controller",
			"version": "1.0-alpha.3",
		})
	})

	handlers.RegisterRoutes(app, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("KingDNS Controller running on :" + port)
	log.Fatal(app.Listen(":" + port))
}
