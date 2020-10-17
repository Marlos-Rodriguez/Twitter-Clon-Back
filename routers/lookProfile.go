package routers

import (
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/gofiber/fiber/v2"
)

//LookProfile permite extraer los valores del perfil
func LookProfile(c *fiber.Ctx) error {

	//Get the id from the URL
	ID := c.Query("id")

	//If the ID is too short
	if len(ID) < 1 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID is required"})
	}

	//Search the user in the DB
	perfil, err := db.SearchUser(ID)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": err})
	}

	return c.JSON(perfil)
}
