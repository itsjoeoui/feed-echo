package handler

import (
	"context"
	"feed/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CreateUser(c echo.Context) error {
	usersColl := h.DB.Collection("users")

	u := &model.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if u.Username == "" || u.DisplayName == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid username, displayname or password"}
	}

	filter := bson.D{{Key: "username", Value: u.Username}}
	count, err := usersColl.CountDocuments(context.TODO(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if count > 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Username already exists"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	u.Password = string(hashedPassword)

	result, err := usersColl.InsertOne(context.TODO(), u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var insertedUser bson.M
	err = usersColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&insertedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Debug(insertedUser)

	return c.JSON(http.StatusOK, insertedUser)
}
