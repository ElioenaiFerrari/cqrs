package entities

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (e *Entity) Prepare() error {
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().Format(time.RFC3339)
	e.UpdatedAt = time.Now().Format(time.RFC3339)

	return nil
}
