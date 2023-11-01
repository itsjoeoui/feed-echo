package handler

import (
	"context"
	"feed/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUser(c echo.Context) error {
	usersColl := h.DB.Collection("users")

	user := &model.User{
		Username:    "jyuhq",
		DisplayName: "Joey Yu",
	}

	result, err := usersColl.InsertOne(context.TODO(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Debug(result)

	return c.JSON(http.StatusOK, user)
}
