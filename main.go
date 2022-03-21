package main

import (
	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
	factories "github.com/ElioenaiFerrari/cqrs/src/factories"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entities.Student{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entities.Course{}); err != nil {
		panic(err)
	}

	nc, err := nats.Connect("nats://localhost:4222")

	if err != nil {
		panic(err)
	}

	e := echo.New()
	r := e.Group("/api/v1")

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	studentController := factories.MakeStudentController(db, nc)

	r.POST("/students", studentController.Create)

	e.Logger.Fatal(e.Start(":3000"))
}
