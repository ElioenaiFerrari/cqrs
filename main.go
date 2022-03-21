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
	db, err := gorm.Open(postgres.Open("host=postgres port=5432 user=postgres dbname=postgres password=postgres sslmode=disable TimeZone=America/Sao_Paulo"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Student{})
	db.AutoMigrate(&entities.Course{})

	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	})

	if err != nil {
		panic(err)
	}
}

func main() {

	e := echo.New()
	r := e.Group("/api/v1")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	studentController := factories.MakeStudentController(db, producer)

	r.POST("/students", studentController.Create)

	e.Logger.Fatal(e.Start(":3000"))
}
