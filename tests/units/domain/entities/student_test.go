package entities_test

import (
	"testing"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
)

func studentFixture() entities.Student {
	return entities.Student{
		RA:    "1234567890",
		Name:  "Elioenai Ferrari",
		Email: "elioenaiferrari@gmail.com",
		Phone: "27999999999",
		CPF:   "12345678901",
		Age:   22,
	}
}

func TestStudentPrepareWhenPassValidParams(t *testing.T) {
	student := studentFixture()

	if err := student.Prepare(); err != nil {
		t.Errorf("Prepare() error = %v, wantErr %v", err, nil)
	}
}

func TestStudentPrepareWhenPassInvalidRA(t *testing.T) {
	student := studentFixture()
	student.RA = "123456789"

	err := student.Prepare()

	if err.Error() != "ra must be 10 digits" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}

func TestStudentPrepareWhenPassInvalidAge(t *testing.T) {
	student := studentFixture()
	student.Age = -1

	err := student.Prepare()

	if err.Error() != "age must be at least 18 years old" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}

func TestStudentPrepareWhenPassInvalidEmail(t *testing.T) {
	student := studentFixture()
	student.Email = "elioeanaiferrari"

	err := student.Prepare()

	if err.Error() != "invalid email" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}

func TestStudentPrepareWhenPassInvalidName(t *testing.T) {
	student := studentFixture()
	student.Name = ""

	err := student.Prepare()

	if err.Error() != "name must be at least 3 characters" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}
