package routers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
)

//ReadTweet Read the tweets
func ReadTweet(c *fiber.Ctx) error {
	//Get The ID
	ID := c.Query("id")

	//Check the query values
	if len(ID) < 1 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "ID is required"})
	}

	if len(c.Query("page")) < 1 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Page param is required"})
	}

	//Converto to int
	pagina, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Page bust be greater than 0", "data": err})
	}

	//Convert to int64
	pag := int64(pagina)

	//Get the tweets from the DB
	response, isFine := db.ReadTweets(ID, pag)

	if isFine == false {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Reading the tweets"})

	}

	//Return the tweets
	return c.JSON(response)
}
