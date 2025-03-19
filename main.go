package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//loading environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "AI Task Manager API is running"})
	})

	app.Post("/register", Register)
	app.Post("/login", Login)

	app.Post("/tasks", CreateTask)
	app.Get("/tasks", GetTasks)

	log.Fatal(app.Listen(":5000"))
}
