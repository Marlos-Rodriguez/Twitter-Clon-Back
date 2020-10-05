package routers

import (
	"errors"
	"os"
	"strings"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email valor de email usado en todos los EndPoints
var Email string

//IDUsuario valor de ID devuelto del modelo, usado en todos los EndPoints
var IDUsuario string

//ProcessToken proceso token para extraer sus valores
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	secrectENV := os.Getenv("SECRECT_KEY")

	myPass := []byte(secrectENV)

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPass, nil
	})

	if err == nil {
		_, encontrado, _ := db.CheckExistingUser(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
