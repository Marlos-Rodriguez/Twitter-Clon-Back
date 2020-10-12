package db

import (
	"context"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertTweet Save the Tweet in DB
func InsertTweet(t models.SaveTweet) (string, bool, error) {
	//Create the context for the DB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//select the collection of database
	db := MongoCN.Database("twitterclon")
	col := db.Collection("tweet")

	//Tweet Model to BSON
	register := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	//Insert in the DB
	result, err := col.InsertOne(ctx, register)

	if err != nil {
		return string(""), false, err
	}

	//Get the ID
	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
