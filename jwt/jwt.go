package jwt

import (
	"os"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT Genera el encriptado con JWT
func GenerateJWT(t models.Usuario) (string, error) {
	//Get the secret password from the ENV
	secrectENV := os.Getenv("SECRECT_KEY")

	myPass := []byte(secrectENV)

	//Make the info for the JWT
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

	//Generates the JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	//Assing the password to JWT
	tokenStr, err := token.SignedString(myPass)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
