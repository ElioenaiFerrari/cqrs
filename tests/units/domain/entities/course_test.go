package entities_test

import (
	"testing"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
)

func courseFixture() entities.Course {
	return entities.Course{
		Name:        "Go for Beginners",
		Description: "Go is a programming language",
		Price:       100.00,
	}
}

func TestCoursePrepareWhenPassValidParams(t *testing.T) {
	course := courseFixture()
	if err := course.Prepare(); err != nil {
		t.Errorf("Prepare() error = %v, wantErr %v", err, nil)
	}
}

func TestCoursePrepareWhenPassInvalidPrice(t *testing.T) {
	course := courseFixture()
	course.Price = -1

	err := course.Prepare()

	if err.Error() != "price must be greater than 0" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}

func TestCoursePrepareWhenPassInvalidName(t *testing.T) {
	course := entities.Course{
		Name:        "",
		Description: "Go is a programming language",
		Price:       100.00,
	}

	err := course.Prepare()

	if err.Error() != "name must be at least 3 characters" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}
