package routers

import (
	"errors"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
	jwt "github.com/dgrijalva/jwt-go"
)

//ProcessToken proceso token para extraer sus valores
func ProcessToken(tk *jwt.Token, UserID string) (string, error) {

	//Assing the claims in a variable
	claims := tk.Claims.(jwt.MapClaims)

	//If UserID it's not empty, verify if the same of JWT
	if UserID != "" && len(UserID) > 1 {
		if claims["_id"].(string) != UserID {
			return "", errors.New("Token ID not match")
		}
	} else {
		UserID = claims["_id"].(string) //If User If UserID it's empty, assing the ID'claims
	}

	//Get the email from the claims
	userEmail := claims["email"].(string)

	//Check if the user Exists in the DB
	_, found, _ := db.CheckExistingUser(userEmail)

	if !found {
		return "", errors.New("User not found with that ID")
	}

	return UserID, nil
}
