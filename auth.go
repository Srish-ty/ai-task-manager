package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("your-secret-key")

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

// Temporary Dummy database
var users = map[string]User{}

func Register(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	user.ID = uuid.NewString()
	users[user.Username] = *user
	return c.JSON(fiber.Map{"message": "User registered"})
}

func Login(c *fiber.Ctx) error {
	data := new(User)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	user, exists := users[data.Username]
	if !exists {
		return fiber.ErrUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": tokenString})
}
