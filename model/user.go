package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	User struct {
		// This is the same 'sub' from Google OAuth
		Sub         string
		DisplayName string
		ID          primitive.ObjectID `bson:"_id,omitempty"`
	}
)
