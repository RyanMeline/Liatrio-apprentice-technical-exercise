package main

import (
	"github.com/gofiber/fiber/v3"
	"time"
)

type Response struct {
	Message		string		`json:"message"`
	Timestamp	int64		`json:"timestamp"`

}

func main() {
	app := fiber.New();

	app.Get("/", func(c fiber.Ctx) error {	
		response := Response {
			Message:	"My name is Ryan Meline",
			Timestamp:	time.Now().Unix(),
		}
		return c.JSON(response)
	})

	app.Listen(":80");
}
