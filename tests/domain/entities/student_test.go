package entities_test

import (
	"testing"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
)

func TestStudentPrepareWhenPassValidParams(t *testing.T) {
	student := entities.Student{
		RA:    "1234567890",
		Name:  "Elioenai Ferrari",
		Email: "elioenaiferrari@gmail.com",
		Phone: "27999999999",
		CPF:   "12345678901",
	}

	if err := student.Prepare(); err != nil {
		t.Errorf("Prepare() error = %v, wantErr %v", err, nil)
	}
}

func TestStudentPrepareWhenPassInvalidRA(t *testing.T) {
	student := entities.Student{
		RA:    "12345678",
		Name:  "Elioenai Ferrari",
		Email: "elioeanaiferrari@gmail.com",
		Phone: "27999999999",
		CPF:   "12345678901",
	}

	err := student.Prepare()

	if err.Error() != "ra must be 10 digits" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}

func TestStudentPrepareWhenPassInvalidEmail(t *testing.T) {
	student := entities.Student{
		RA:    "12345678",
		Name:  "Elioenai Ferrari",
		Email: "elioeanaiferrari",
		Phone: "27999999999",
		CPF:   "12345678901",
	}

	err := student.Prepare()

	if err.Error() != "invalid email" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}

func TestStudentPrepareWhenPassInvalidName(t *testing.T) {
	student := entities.Student{
		RA:    "1234567890",
		Name:  "",
		Email: "elioeanaiferrari@gmail.com",
		Phone: "27999999999",
		CPF:   "12345678901",
	}

	err := student.Prepare()

	if err.Error() != "name must be at least 3 characters" {
		t.Errorf("Prepare() error = %v, wantErr %v", err, "invalid email")
	}
}
