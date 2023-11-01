package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Post struct {
		Content string
		ID      primitive.ObjectID `bson:"_id,omitempty"`
		Author  primitive.ObjectID
		Likes   int
	}
)
