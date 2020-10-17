package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
)

//ChequeoDB Middleware que comprueba la conexion de la DB
func ChequeoDB() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !db.CheckConnection() {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   "Error Connect in the DB",
			})
		}
		return c.Next()
	}
}
