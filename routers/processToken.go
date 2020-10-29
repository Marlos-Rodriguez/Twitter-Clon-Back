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
func ProcessToken(tk *jwt.Token) error {

	//Assing the claims in a variable
	claims := tk.Claims.(jwt.MapClaims)

	UserID = claims["_id"].(string)
	UserEmail = claims["email"].(string)

	//Check if the user Exists in the DB
	_, found, _ := db.CheckExistingUser(UserEmail)

	if !found {
		return errors.New("User not found with that ID")
	}

	return nil
}
