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
	//Get the secret key from ENV
	secrectENV := os.Getenv("SECRECT_KEY")

	//Convert to bytes
	myPass := []byte(secrectENV)

	claims := &models.Claim{}

	//Separate Bearer of the token
	splitToken := strings.Split(tk, "Bearer")

	//If the token not a two map object
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	//Eliminates empty spaces of token
	tk = strings.TrimSpace(splitToken[1])

	//Decrypt the token with the password
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPass, nil
	})

	//If everything was fine with the token
	if err == nil {
		//Search the user with the info from the token
		_, encontrado, _ := db.CheckExistingUser(claims.Email)
		//Assing the info to the claim
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
