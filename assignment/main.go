package main

import (
	"Desktop/shopi/assignment/handler"
	"Desktop/shopi/assignment/repository"
	"Desktop/shopi/assignment/service"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	app := NewApplication()
	app.Listen(":8080")

}

func NewApplication() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH",
	}))
	app.Use(logger.New())
	// Database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://aykutisik:Ayk-0109@firstcluster.o4wm6.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)

	}
	database := client.Database("myFirstDatabase")
	thecollection := database.Collection("orders2")
	repo := repository.NewRepository(database, client, thecollection)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	app.Post("/Add", handler.Add)
	app.Post("/Filter", handler.Filter)

	//

	return app
}
