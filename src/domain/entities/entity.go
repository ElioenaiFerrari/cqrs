package entities

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (e *Entity) Prepare() error {
	e.ID = uuid.New()
	e.CreatedAt = time.Now().UTC()
	e.UpdatedAt = time.Now().UTC()

	return nil
}
