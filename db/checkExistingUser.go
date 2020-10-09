package db

import (
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckExistingUser recibe un email y verifica si ya existe en la DB
func CheckExistingUser(email string) (models.Usuario, bool, string) {

	//Make Chanel for gorutine with Context & Collection mongo method
	c := make(chan models.ContextModel)

	//Gorutine for Get Context & Collection mongo method
	go CreateContext("users", c)

	condition := bson.M{"email": email}

	var resultado models.Usuario

	cntxt := <-c

	//Searh for a user with same email
	err := cntxt.Col.FindOne(cntxt.Ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()

	defer cntxt.Cancel()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
