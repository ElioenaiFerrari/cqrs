package main

import (
	"context"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
	factories "github.com/ElioenaiFerrari/cqrs/src/factories"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	writeDB, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := writeDB.AutoMigrate(&entities.Student{}); err != nil {
		panic(err)
	}

	if err := writeDB.AutoMigrate(&entities.Course{}); err != nil {
		panic(err)
	}

	mongo, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	if err := mongo.Connect(context.TODO()); err != nil {
		panic(err)
	}

	readDB := mongo.Database("cqrs")

	nc, err := nats.Connect("nats://localhost:4222")

	if err != nil {
		panic(err)
	}

	e := echo.New()
	r := e.Group("/api/v1")

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	studentController := factories.MakeStudentController(writeDB, readDB, nc)

	r.POST("/students", studentController.Create)
	r.GET("/students", studentController.Index)

	e.Logger.Fatal(e.Start(":3000"))
}
