package routers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//SaveTwitter Save the tweet in the DB
func SaveTwitter(c *fiber.Ctx) error {
	//Basic model for decode body
	var message models.Tweet

	//Decode the body of request
	if err := c.BodyParser(&message); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	//Get the Token
	tk := c.Locals("user").(*jwt.Token)

	//If User exists, verify if ID mach or return the id from the Claims
	IDUser, err := ProcessToken(tk, "")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error proccesing the token", "data": err.Error()})
	}

	//Create the complete tweet model
	register := models.SaveTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	//Insert in the DB
	_, status, err := db.InsertTweet(register)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error creating the Tweet in the DB", "data": err.Error()})
	}

	if status == false {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "The tweet is not inserted"})
	}

	//Return a 200 code if everything is fine
	return c.Status(200).JSON(fiber.Map{"status": "successful", "message": "Tweet Created"})
}
