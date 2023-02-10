package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1552811",
		Key:     "0c5793d36c7a6d98b9fb",
		Secret:  "8a0939704af3a9495a75",
		Cluster: "ap1",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {

		var data map[string]string
		err := c.BodyParser(&data)
		if err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)

		return c.SendString("Succes")
	})

	app.Listen(":9090")
}
