package routers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//SaveRelation Save relation between users
func SaveRelation(c *fiber.Ctx) error {
	ID := c.Query("id")

	if len(ID) < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "ID is required"})
	}

	var t models.Relation

	tk := c.Locals("user").(*jwt.Token)

	if err := ProcessToken(tk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error processing token", "data": err.Error()})
	}

	t.UserID = UserID
	t.UserRelationID = ID

	if status, err := db.InsertRelation(t); err != nil || !status {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error Inserting in the DB", "data": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}
