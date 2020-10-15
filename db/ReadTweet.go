package db

import (
	"context"
	"log"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadTweets Get the tweets of Users
func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	//Make Chanel for gorutine with Context & Collection mongo method
	c := make(chan models.ContextModel)

	//Gorutine for Get Context & Collection mongo method
	go CreateContext("tweet", c)

	var results []*models.ReturnTweets

	condition := bson.M{
		"userid": ID,
	}

	optionsDB := options.Find()

	optionsDB.SetLimit(20)
	optionsDB.SetSort(bson.D{{Key: "fecha", Value: -1}})
	optionsDB.SetSkip((page - 1) * 20)

	cntxt := <-c

	cursor, err := cntxt.Col.Find(cntxt.Ctx, condition, optionsDB)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.ReturnTweets

		err := cursor.Decode(&register)

		if err != nil {
			return results, false
		}

		results = append(results, &register)
	}

	return results, true
}
