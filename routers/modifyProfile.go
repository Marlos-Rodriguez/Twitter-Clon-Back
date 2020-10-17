package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dgrijalva/jwt-go"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//ModifyProfile modifica el perfil de usuario
func ModifyProfile(c *fiber.Ctx) error {
	//Create a base User Model
	var t models.Usuario

	//Decode the body of request
	if err := c.BodyParser(&t); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	var status bool

	//Get the JWT
	tk := c.Locals("user").(*jwt.Token)

	//Verify if user exits, And is the same ID or get the UserID from the Token
	IDUser, err := ProcessToken(tk, "")

	//If any error in the token
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error proccesing the token", "data": err.Error()})
	}

	//Try to modify the profile in DB
	status, err = db.ModifyProfile(t, IDUser)

	//If errors in the DB
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error in modify Profile in DB", "data": err.Error()})
	}

	if !status {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Wrog user Data"})

	}

	//Return a 200 code if everything is fine
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": "successful", "message": "Modified profile"})
}
