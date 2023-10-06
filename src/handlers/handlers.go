package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/materdev/golang_test-server/src/dice"
)

// Basic health check
func HealthCheck(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

// Handler to roll a full set of dice
	// RollAllDice uses a struct in dice/rollDice.go
func RollAllDiceHandler(c *fiber.Ctx) error {
    rolls := dice.RollAllDice()  // Call the utility function
    rollsJSON, _ := json.Marshal(rolls)  // Convert the result to JSON
    return c.Send(rollsJSON)  // Send the result back as a JSON response
}

// Struct for single dice responses
type DiceRollResponse struct {
	DiceType string `json:"diceType"`
	Value    int    `json:"value"`
}

// Individual dice roll handlers
func RollD4Handler(c *fiber.Ctx) error {
	roll := dice.RollD4()
	rollResult := DiceRollResponse{"D4", roll}
	rollJSON, _ := json.Marshal(rollResult)
	return c.Send(rollJSON)
}

func RollD6Handler(c *fiber.Ctx) error {
	roll := dice.RollD6()
	rollResult := DiceRollResponse{"D6", roll}
	rollJSON, _ := json.Marshal(rollResult)
	return c.Send(rollJSON)
}

func RollD8Handler(c *fiber.Ctx) error {
	roll := dice.RollD8()
	rollResult := DiceRollResponse{"D8", roll}
	rollJSON, _ := json.Marshal(rollResult)
	return c.Send(rollJSON)
}

func RollD10Handler(c *fiber.Ctx) error {
	roll := dice.RollD10()
	rollResult := DiceRollResponse{"D10", roll}
	rollJSON, _ := json.Marshal(rollResult)
	return c.Send(rollJSON)
}

func RollD100Handler(c *fiber.Ctx) error {
	roll := dice.RollD100()
	rollResult := DiceRollResponse{"D100", roll}
	rollJSON, _ := json.Marshal(rollResult)
	return c.Send(rollJSON)
}

func RollD12Handler(c *fiber.Ctx) error {
	roll := dice.RollD12()
	rollResult := DiceRollResponse{"D12", roll}
	rollJSON, _ := json.Marshal(rollResult)
	return c.Send(rollJSON)
}

func RollD20Handler(c *fiber.Ctx) error {
	roll := dice.RollD20()
	rollResult := DiceRollResponse{"D20", roll}
	rollJSON, _ := json.Marshal(rollResult)
	return c.Send(rollJSON)
}
