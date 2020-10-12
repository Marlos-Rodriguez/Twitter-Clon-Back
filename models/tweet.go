package models

//Tweet Model
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
