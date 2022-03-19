package entities

import (
	"errors"
	"net/mail"
)

type Student struct {
	Entity
	RA    string `json:"ra"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	CPF   string `json:"cpf"`
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
		return errors.New("invalid phone")
	}

	if len(s.CPF) != 11 {
		return errors.New("invalid cpf")
	}

	if len(s.RA) != 10 {
		return errors.New("invalid ra")
	}

	if len(s.Name) < 3 {
		return errors.New("invalid name")
	}

	return nil
}
