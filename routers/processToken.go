package routers

import (
	"errors"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	jwt "github.com/dgrijalva/jwt-go"
)

//UserID ID of user from Claims
var UserID string

//UserEmail Email of User from Claims
var UserEmail string

//ProcessToken proceso token para extraer sus valores
func ProcessToken(tk *jwt.Token, IDUser string) error {

	//Assing the claims in a variable
	claims := tk.Claims.(jwt.MapClaims)

	//If UserID it's not empty, verify if the same of JWT
	if IDUser != "" && len(IDUser) > 1 {
		if claims["_id"].(string) != IDUser {
			return errors.New("Token ID not match")
		}
	}

	if len(UserID) < 1 && len(UserEmail) < 1 {
		UserID = claims["_id"].(string)
		UserEmail = claims["email"].(string)
	}

	//Check if the user Exists in the DB
	_, found, _ := db.CheckExistingUser(UserEmail)

	if !found {
		return errors.New("User not found with that ID")
	}

	return nil
}
