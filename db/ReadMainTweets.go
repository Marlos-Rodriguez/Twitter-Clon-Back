package db

import (
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ReadMainTweets Return all tweets of main page
func ReadMainTweets(ID string, page int) ([]*models.ReturnMainTweets, bool) {
	//Make Chanel for gorutine with Context & Collection mongo method
	c := make(chan models.ContextModel)

	//Gorutine for Get Context & Collection mongo method
	go CreateContext("relation", c)

	skip := (page - 1) * 20

	//Array of conditions
	conditions := make([]bson.M, 0)

	//Get all relations with the User ID
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})

	//Relation all the tweets of all the person who user follow
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})

	//Destruct the info to read it
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	//Sort for date
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	//Skip for the number of page
	conditions = append(conditions, bson.M{"$skip": skip})
	//limit for page
	conditions = append(conditions, bson.M{"$limit": 20})

	cntxt := <-c

	defer cntxt.Cancel()

	cursor, err := cntxt.Col.Aggregate(cntxt.Ctx, conditions)

	var result []*models.ReturnMainTweets

	if err = cursor.All(cntxt.Ctx, &result); err != nil {
		return result, false
	}

	return result, true
}
