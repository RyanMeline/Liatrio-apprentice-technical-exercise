package main

import (
	"github.com/gofiber/fiber/v3"
	"time"
)

type Response struct {
	Message		string		`json:"message"`
	Timestamp	time.Time	`json:"timestamp"`

}

func main() {
	app := fiber.New();

	response := Response {
		Message:	"My name is Ryan Meline",
		Timestamp:	time.Now(),
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(response)
	})

	app.Listen(":3000");
}
