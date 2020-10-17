package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//Registro funcion para crear nuevo usuario en DB
func Registro(c *fiber.Ctx) error {
	//Create a base User model
	var t models.Usuario

	//Decode the body of request
	if err := c.BodyParser(&t); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	if len(t.Email) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Email is required"})

	}
	if len(t.Password) < 6 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Password bust be greater than 0"})
	}

	//Search if the user is alredy exits
	_, encontrado, _ := db.CheckExistingUser(t.Email)

	if encontrado {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User already exists"})
	}

	//if it doesn't exist yet,create it
	_, status, err := db.InsertRegistro(t)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error creating the user in the DB", "data": err.Error()})
	}

	if status == false {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "The User is not inserted", "data": err.Error()})
	}

	//Return a 200 code if everything is fine
	return c.Status(200).JSON(fiber.Map{"status": "successful", "message": "Modified profile"})
}
