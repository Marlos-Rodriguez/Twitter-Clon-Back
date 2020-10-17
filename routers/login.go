package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/jwt"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//Login realiza el login
func Login(c *fiber.Ctx) error {

	//Create base User model
	var t models.Usuario

	//Decode the body of request
	if err := c.BodyParser(&t); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	//If the email of the body is Empty
	if len(t.Email) == 0 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Email is required"})
	}

	//Try the login in the DB
	documento, exist := db.IntentoLogin(t.Email, t.Password)

	//If user not exist
	if !exist {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "User Not exits"})
	}

	//Generate a JWT
	jwtKey, err := jwt.GenerateJWT(documento)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error in generate the token", "data": err})
	}

	//Return the JWT
	return c.JSON(&fiber.Map{
		"success": true,
		"Token":   jwtKey,
	})
}
