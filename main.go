package main

import (
	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
	factories "github.com/ElioenaiFerrari/cqrs/src/factories"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var producer *kafka.Producer

func init() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=cqrs password=postgres sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Student{})

	producer, _ = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
}

func main() {

	e := echo.New()
	r := e.Group("/api/v1")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	studentController := factories.MakeStudentController(db, producer)

	r.POST("/student", studentController.Create)

	e.Logger.Fatal(e.Start(":3000"))
}
