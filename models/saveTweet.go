package models

import "time"

//SaveTweet Tweet Model
type SaveTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
