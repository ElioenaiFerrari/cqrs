package entities

import (
	"time"
)

type Entity struct {
	ID        uint64    `json:"id" gorm:"primaryKey" bson:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
