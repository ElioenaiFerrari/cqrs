package entities

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (e *Entity) Prepare() error {
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().UTC()
	e.UpdatedAt = time.Now().UTC()

	return nil
}
