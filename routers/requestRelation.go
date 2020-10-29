package routers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//RequestRelation Check if A relation between users exists
func RequestRelation(c *fiber.Ctx) error {
	//Get The ID
	ID := c.Query("id")

	//Check the query values
	if len(ID) < 1 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID is required"})
	}

	var t models.Relation

	tk := c.Locals("user").(*jwt.Token)

	if err := ProcessToken(tk); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error proccesing the token", "data": err.Error()})
	}

	t.UserID = UserID
	t.UserRelationID = ID

	var resp models.ResponseRequestRelation

	if status, err := db.RequestRelation(t); err != nil || !status {
		resp.Status = false
	}

	resp.Status = true

	return c.Status(fiber.StatusAccepted).JSON(resp)
}
