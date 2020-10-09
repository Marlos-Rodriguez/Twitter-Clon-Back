package db

import (
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertRegistro funcion para insertar nuevo usuario en DB
func InsertRegistro(u models.Usuario) (string, bool, error) {
	//Make chanel for Context & Collection mongo method
	c := make(chan models.ContextModel)

	//Gorutine for Get Context & Collection mongo method
	go CreateContext("users", c)

	//Encrypt the Password
	u.Password, _ = EncryptPassword(u.Password)

	cntxt := <-c

	//Insert User in Database
	result, err := cntxt.Col.InsertOne(cntxt.Ctx, u)

	defer cntxt.Cancel()

	if err != nil {
		return "", false, err
	}

	//Insert a ID in the User Object
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
