package routers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/dgrijalva/jwt-go"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
)

//ListUsers Read List of Users
func ListUsers(c *fiber.Ctx) error {
	typeUser := c.Query("type")
	page := c.Query("page")
	search := c.Query("search")

	pagTemp, err := strconv.Atoi(page)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Page must be bigger that 0", "data": err.Error()})
	}

	pag := int64(pagTemp)

	//Get the JWT
	tk := c.Locals("user").(*jwt.Token)

	//Verify if user exits, And is the same ID or get the UserID from the Token
	if err := ProcessToken(tk); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error proccesing the token", "data": err.Error()})
	}

	result, status := db.ReadAllUsers(UserID, pag, search, typeUser)

	if status == false {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error proccesing the Users"})
	}

	return c.Status(fiber.StatusAccepted).JSON(result)
}
