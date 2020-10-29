package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ReturnMainTweets Main tweets struct
type ReturnMainTweets struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userrelationid" json:"userRelationId,omitempty"`
	UserRelationID string             `bson:"userid" json:"userId,omitempty"`
	Tweet          maintweet
}

type maintweet struct {
	TweetID string    `bson:"tweetid" json:"TweetId,omitempty"`
	Message string    `bson:"message" json:"Message,omitempty"`
	Date    time.Time `bson:"date" json:"Date,omitempty"`
}
