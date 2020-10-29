package routers

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
)

//ReadMainTweets return the tweets of main page
func ReadMainTweets(c *fiber.Ctx) error {
	if len(c.Query("page")) < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Must be send page parameter"})
	}

	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Must be send page parameter like a integer", "data": err.Error()})
	}

	tk := c.Locals("user").(*jwt.Token)

	if err := ProcessToken(tk, ""); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error proccesing the token", "data": err.Error()})
	}

	response, isFine := db.ReadMainTweets(UserID, page)

	if !isFine {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Requesting the tweets from DB", "data": err.Error()})
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}
