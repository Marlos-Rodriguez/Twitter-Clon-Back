package db

import (
	"context"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//InsertRelation Save relation in DB
func InsertRelation(t models.Relation) (bool, error) {
	//Create the context for the DB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//select the collection of database
	db := MongoCN.Database("twitterclon")
	col := db.Collection("relation")

	if _, err := col.InsertOne(ctx, t); err != nil {
		return false, err
	}

	return true, nil
}
