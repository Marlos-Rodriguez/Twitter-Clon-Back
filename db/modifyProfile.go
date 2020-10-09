package db

import (
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModifyProfile modifica el perfil en la DB
func ModifyProfile(u models.Usuario, ID string) (bool, error) {
	//Make Chanel for gorutine with Context & Collection mongo method
	c := make(chan models.ContextModel)

	//Gorutine for Get Context & Collection mongo method
	go CreateContext("users", c)

	//Make a base register interface for modify user
	registro := make(map[string]interface{})

	//Verify all values of User and assign the existing ones
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}

	registro["fechaNacimiento"] = u.Fecha

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}

	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}

	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	//Convert to bson native
	updtString := bson.M{
		"$set": registro,
	}

	//Converte ID to Object ID
	objID, _ := primitive.ObjectIDFromHex(ID)

	//Create filter for the DB
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	cntxt := <-c

	//Update the user if is found
	_, err := cntxt.Col.UpdateOne(cntxt.Ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	cntxt.Cancel()

	return true, nil
}
