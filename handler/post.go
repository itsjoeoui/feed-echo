package handler

import (
	"context"
	"feed/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) CreatePost(c echo.Context) error {
	// TODO: remove the hardcoded user later
	usersColl := h.DB.Collection("users")

	u := &model.User{}
	filter := bson.D{{Key: "username", Value: "jyuhq"}}
	err := usersColl.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Create a new Post and load the request body into it
	p := &model.Post{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Make sure we get all the fields we need
	if p.Content == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Post content cannot be empty!"}
	}
	p.Author = u.ID
	p.Date = time.Now()

	postsColl := h.DB.Collection("posts")

	// Actually insert the new user to DB
	result, err := postsColl.InsertOne(context.TODO(), p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return the inserted user
	var insertedPost bson.M
	err = postsColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&insertedPost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, insertedPost)
}

func (h *Handler) GetPosts(c echo.Context) error {
	// TODO: remove the hardcoded user later
	usersColl := h.DB.Collection("users")

	u := &model.User{}
	filter := bson.D{{Key: "username", Value: "jyuhq"}}
	err := usersColl.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	postsColl := h.DB.Collection("posts")

	// Actually insert the new user to DB
	cursor, err := postsColl.Find(context.TODO(), bson.D{{Key: "author", Value: u.ID}})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer cursor.Close(context.TODO())

	var posts []model.Post

	if err := cursor.All(context.TODO(), &posts); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, posts)
}
