package db

import (
	"context"
	"log"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SearchUser busca datos de un usuario en la DB
func SearchUser(ID string) (models.Usuario, error) {
	//Create the context for the DB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//select the collection of database
	db := MongoCN.Database("twitterclon")
	col := db.Collection("users")

	//Create base model User
	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	//Find a user in the DB
	if err := col.FindOne(ctx, condition).Decode(&perfil); err != nil {
		//Make empty the password
		perfil.Password = ""
		log.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}

	//Make empty the password
	perfil.Password = ""

	return perfil, nil
}
