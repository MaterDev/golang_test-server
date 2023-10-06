package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/materdev/golang_test-server/src/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/healthCheck", handlers.HealthCheck)

	// Roll full set of dice
	app.Get("/roll/all", handlers.RollAllDiceHandler)

    // Individual dice roll endpoints
    app.Get("/roll/d4", handlers.RollD4Handler)
    app.Get("/roll/d6", handlers.RollD6Handler)
    app.Get("/roll/d8", handlers.RollD8Handler)
    app.Get("/roll/d10", handlers.RollD10Handler)
    app.Get("/roll/d100", handlers.RollD100Handler)
    app.Get("/roll/d12", handlers.RollD12Handler)
    app.Get("/roll/d20", handlers.RollD20Handler)

	app.Listen(":3000")
}