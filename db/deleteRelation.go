package db

import (
	"context"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//DeleteRelation Borra la relacion con la DB
func DeleteRelation(t models.Relation) (bool, error) {
	//Create the context for the DB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//select the collection of database
	db := MongoCN.Database("twitterclon")
	col := db.Collection("relation")

	if _, err := col.DeleteOne(ctx, t); err != nil {
		return false, err
	}

	return true, nil
}
