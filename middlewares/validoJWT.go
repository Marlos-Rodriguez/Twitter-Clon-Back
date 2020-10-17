package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v2"
)

//ValidJWT permite validar el JWT que nos viene en la peticion
func ValidJWT() fiber.Handler {
	SECRETWORD := os.Getenv("SECRECT_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(SECRETWORD),
	})
}
