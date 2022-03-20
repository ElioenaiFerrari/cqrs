package entities_test

import (
	"testing"
	"time"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
)

func TestEntityPrepare(t *testing.T) {
	entity := entities.Entity{}

	if err := entity.Prepare(); err != nil {
		t.Errorf("Prepare() error = %v, wantErr %v", err, nil)
	}

	if entity.ID == "" {
		t.Errorf("Prepare() error = %v, wantErr %v", entity.ID, "")
	}

	if entity.CreatedAt == (time.Time{}) {
		t.Errorf("Prepare() error = %v, wantErr %v", entity.CreatedAt, "")
	}

	if entity.UpdatedAt == (time.Time{}) {
		t.Errorf("Prepare() error = %v, wantErr %v", entity.UpdatedAt, "")
	}

}
