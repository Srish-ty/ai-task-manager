package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var tasks = []Task{}

func CreateTask(c *fiber.Ctx) error {
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return err
	}

	task.ID = uuid.NewString()
	task.Status = "pending"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	tasks = append(tasks, *task)
	return c.JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	return c.JSON(tasks)
}
