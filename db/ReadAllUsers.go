package db

import (
	"context"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadAllUsers Read the register users and returns only the related ones
func ReadAllUsers(ID string, page int64, search string, Stype string) ([]*models.Usuario, bool) {
	//Create the context for the DB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//select the collection of database
	db := MongoCN.Database("twitterclon")
	col := db.Collection("users")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		return results, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var s models.Usuario

		err := cur.Decode(&s)

		if err != nil {
			return results, false
		}

		var r models.Relation

		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, err = RequestRelation(r)

		if Stype == "new" && !found {
			include = true
		}

		if Stype == "follow" && found {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		if include {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}

	cur.Close(ctx)

	return results, true
}
