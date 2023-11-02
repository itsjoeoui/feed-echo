package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Post struct {
		Date    time.Time
		Content string
		ID      primitive.ObjectID `bson:"_id,omitempty"`
		Author  primitive.ObjectID
		Likes   int
	}
)
