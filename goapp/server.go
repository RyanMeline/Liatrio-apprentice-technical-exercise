package main

import (
	"github.com/gofiber/fiber/v3"
	"time"
	"encoding/json"
)

type Response struct {
	Message		string		`json:"message"`
	Timestamp	int64		`json:"timestamp"`
	Test 		string		`json:"test"`
}

func main() {
	app := fiber.New(fiber.Config{ JSONEncoder: json.Marshal, });

	app.Get("/", func(c fiber.Ctx) error {	
		response := Response {
			Message:	"My name is Ryan Meline",
			Timestamp:	time.Now().UnixMilli(),
			Test:		"Test string a",
		}
		return c.JSON(response)
	})

	app.Listen(":8080");
}
