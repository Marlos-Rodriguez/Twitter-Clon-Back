package routers

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//UploadAvatar Upload the avatar to DB
func UploadAvatar(c *fiber.Ctx) error {
	handler, err := c.FormFile("avatar")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error in get the image", "data": err.Error()})
	}

	var extension = strings.Split(handler.Filename, ".")[1]

	tk := c.Locals("user").(*jwt.Token)

	//Verify if user exits, And is the same ID or get the UserID from the Token
	if err := ProcessToken(tk, ""); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error proccesing the token", "data": err.Error()})
	}

	var path string = "./uploads/avatars/" + UserID + "." + extension

	if err = c.SaveFile(handler, path); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error in save the image", "data": err.Error()})
	}

	var user models.Usuario

	user.Avatar = UserID + "." + extension

	if status, err := db.ModifyProfile(user, UserID); err != nil || status == false {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Uploading the image", "data": err.Error()})
	}

	return c.SendStatus(fiber.StatusAccepted)
}
