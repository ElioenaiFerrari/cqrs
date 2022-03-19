package entities

import "errors"

type Course struct {
	Entity
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (c Course) Prepare() error {

	if err := c.validate(); err != nil {
		return err
	}

	if err := c.Entity.Prepare(); err != nil {
		return err
	}

	return nil
}

func (c Course) validate() error {
	if len(c.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	if c.Price < 0 {
		return errors.New("price must be greater than 0")
	}

	return nil
}
