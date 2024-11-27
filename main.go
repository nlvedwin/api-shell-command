package main

import (
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"

	"encoding/json"
)

type Convert struct {
	Command string `json:"command"`
}

func main() {
	app := fiber.New()

	app.Post("/convert", func(c *fiber.Ctx) error {
		start := time.Now()
		data := new(Convert)
		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		jsonData, _ := json.Marshal(data)
		println("Received data:", string(jsonData))

		cmd := exec.Command("sh", "-c", data.Command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": string(output)})
		}
		println("output:", string(output))

		executionTime := time.Since(start)
		return c.JSON(fiber.Map{
			"data":          data,
			"executionTime": executionTime.String(),
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Listen(":4001")
}
