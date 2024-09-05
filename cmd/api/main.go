package main

import (
	"github.com/SamuelBarbosa1/ticket-booking-project-v1/config"
	"github.com/SamuelBarbosa1/ticket-booking-project-v1/db"
	"github.com/SamuelBarbosa1/ticket-booking-project-v1/handlers"
	"github.com/SamuelBarbosa1/ticket-booking-project-v1/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	EventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), EventRepository)

	app.Listen(":3000")
}
