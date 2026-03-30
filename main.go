package main

import (
	"stability-test-task-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Task Manager API is running")
	})
	app.Get("/tasks", handlers.GetTasks)
	app.Get("/tasks/:id", handlers.GetTask)
	app.Post("/tasks", handlers.CreateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	app.Listen(":3000")
}
