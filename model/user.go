package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	User struct {
		Username    string
		DisplayName string
		Password    string
		ID          primitive.ObjectID `bson:"_id,omitempty"`
	}
)
