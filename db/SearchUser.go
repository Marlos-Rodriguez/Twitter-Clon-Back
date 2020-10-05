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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterclon")
	col := db.Collection("users")

	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&perfil)

	perfil.Password = ""

	if err != nil {
		log.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}

	return perfil, nil
}
