package jwt

import (
	"os"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT Genera el encriptado con JWT
func GenerateJWT(t models.Usuario) (string, error) {
	secrectENV := os.Getenv("SECRECT_KEY")

	myPass := []byte(secrectENV)

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.Fecha,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(myPass)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
