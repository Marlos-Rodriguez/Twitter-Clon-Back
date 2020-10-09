package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//ContextModel Struct for work with DB
type ContextModel struct {
	Ctx    context.Context
	Col    *mongo.Collection
	Cancel context.CancelFunc
}
