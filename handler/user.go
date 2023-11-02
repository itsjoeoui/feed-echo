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
	// Create a new User and load the request body into it
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Make sure we get all the fields we need
	if u.Username == "" || u.DisplayName == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid username, displayname or password"}
	}

	usersColl := h.DB.Collection("users")
	filter := bson.D{{Key: "username", Value: u.Username}}
	count, err := usersColl.CountDocuments(context.TODO(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if count > 0 {
		return &echo.HTTPError{Code: http.StatusConflict, Message: "Username already exists"}
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	u.Password = string(hashedPassword)

	// Actually insert the new user to DB
	result, err := usersColl.InsertOne(context.TODO(), u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return the inserted user
	var insertedUser bson.M
	err = usersColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&insertedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, insertedUser)
}
