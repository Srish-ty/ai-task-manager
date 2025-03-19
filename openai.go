package main

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	openai "github.com/sashabaranov/go-openai"
)

var openaiClient = openai.NewClient(os.Getenv("OPENAI_API_KEY"))

func SuggestTask(c *fiber.Ctx) error {
	userInput := c.Query("query")
	resp, err := openaiClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-4",
		Messages: []openai.ChatCompletionMessage{
			{Role: "user", Content: "Suggest a task related to: " + userInput},
		},
	})

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"suggestion": resp.Choices[0].Message.Content})
}
