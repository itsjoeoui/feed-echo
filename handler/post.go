package handler

import (
	"context"
	"feed/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) CreatePost(c echo.Context) error {
	usersColl := h.DB.Collection("users")

	var user model.User
	filter := bson.D{{Key: "username", Value: "jyuhq"}}
	err := usersColl.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		c.Logger().Fatal(err)

		return c.JSON(http.StatusInternalServerError, err)
	}

	p := &model.Post{
		Content: "Hello World!",
		Author:  user.ID,
		Likes:   0,
	}

	postsColl := h.DB.Collection("posts")

	result, err := postsColl.InsertOne(context.TODO(), p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Debug(result)

	return c.JSON(http.StatusOK, p)
}

func (h *Handler) GetPost(c echo.Context) error {
	return nil
}
