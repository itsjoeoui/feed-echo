package handler

import "go.mongodb.org/mongo-driver/mongo"

type (
	Handler struct {
		DB *mongo.Session
	}
)
