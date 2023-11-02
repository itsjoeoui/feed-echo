package main

import (
	"context"
	"feed/handler"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)

	if err := godotenv.Load(); err != nil {
		e.Logger.Fatal(err)
	}

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URL")).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		e.Logger.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("feed").RunCommand(
		context.TODO(), bson.D{{Key: "ping", Value: 1}},
	).Err(); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Debug("Pinged your deployment. You successfully connected to MongoDB!")

	h := &handler.Handler{
		DB: client.Database("feed"),
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/post", h.CreatePost)
	e.GET("/post", h.GetPosts)
	e.POST("/user", h.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
