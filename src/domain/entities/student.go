package entities

import (
	"errors"
	"net/mail"
)

type Student struct {
	Entity
	Course string `json:"course"`
	RA     string `json:"ra"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	CPF    string `json:"cpf"`
	Age    int    `json:"age"`
}

func (s Student) Prepare() error {

	if err := s.validate(); err != nil {
		return err
	}

	if err := s.Entity.Prepare(); err != nil {
		return err
	}

	return nil
}

func (s Student) validate() error {
	if _, err := mail.ParseAddress(s.Email); err != nil {
		return errors.New("invalid email")
	}

	if len(s.Phone) != 11 {
		return errors.New("phone must be 11 digits")
	}

	if len(s.CPF) != 11 {
		return errors.New("cpf must be 11 digits")
	}

	if len(s.RA) != 10 {
		return errors.New("ra must be 10 digits")
	}

	if len(s.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	if s.Age < 18 {
		return errors.New("age must be at least 18 years old")
	}

	return nil
}
