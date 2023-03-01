package handlers

import "github.com/gofiber/fiber/v2"

func RoomCreate(c *fiber.Ctx) error {
	return c.JSON("room create")
}
