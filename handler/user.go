package handler

import (
	"feed/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateUser(c echo.Context) error {
	// Create a new User and load the request body into it
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Make sure we get all the fields we need
	if u.DisplayName == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Please provide a valid display name"}
	}

	return nil
}
