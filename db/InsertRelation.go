package db

import (
	"context"
	"errors"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"go.mongodb.org/mongo-driver/bson"
)

//InsertRelation Save relation in DB
func InsertRelation(t models.Relation) (bool, error) {
	//Create the context for the DB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//select the collection of database
	db := MongoCN.Database("twitterclon")
	col := db.Collection("relation")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation

	if err := col.FindOne(ctx, condition).Decode(&result); err == nil {
		return false, errors.New("already existing relation")
	}

	if _, err := col.InsertOne(ctx, t); err != nil {
		return false, err
	}

	return true, nil
}
