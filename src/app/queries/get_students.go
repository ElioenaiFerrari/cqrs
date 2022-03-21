package queries

import (
	"context"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetStudentsQuery struct {
	db *mongo.Database
}

func NewGetStudentsQuery(db *mongo.Database) GetStudentsQuery {
	return GetStudentsQuery{
		db: db,
	}
}

func (gsq *GetStudentsQuery) Execute() ([]entities.Student, error) {
	var students []entities.Student
	collection := gsq.db.Collection("students")

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return students, err
	}

	for cursor.Next(context.TODO()) {
		var student entities.Student
		cursor.Decode(&student)

		students = append(students, student)
	}

	return students, nil
}
