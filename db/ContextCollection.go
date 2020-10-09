package db

import (
	"context"
	"time"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
)

//CreateContext create context for work with DB
func CreateContext(collection string, c chan models.ContextModel) {
	//Create the context for the DB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//select the collection of database
	db := MongoCN.Database("twitterclon")
	col := db.Collection(collection)

	//Create new model for assign the values
	base := new(models.ContextModel)

	//assign the values
	base.Ctx = ctx
	base.Col = col
	base.Cancel = cancel

	//Assign the base model to the Chanel
	c <- *base
}
