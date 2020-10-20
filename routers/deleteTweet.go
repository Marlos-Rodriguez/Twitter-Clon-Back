package routers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
)

//DeleteTweet Delete a tweet for ID
func DeleteTweet(c *fiber.Ctx) error {
	//Get The ID
	ID := c.Query("id")

	//Check the query values
	if len(ID) < 1 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID is required"})
	}

	tk := c.Locals("user").(*jwt.Token)

	UserID := tk.Claims.(jwt.MapClaims)["_id"].(string)

	if len(UserID) < 1 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error in Get the ID from JWT"})
	}

	if err := db.DeleteTweet(ID, UserID); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Deleting the Tweet in the DB", "data": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}
