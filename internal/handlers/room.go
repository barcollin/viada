package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	guid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guid.New().String()))
}
