package entities_test

import (
	"testing"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
)

func TestEntityPrepare(t *testing.T) {
	entity := entities.Entity{}

	if err := entity.Prepare(); err != nil {
		t.Errorf("Prepare() error = %v, wantErr %v", err, nil)
	}
}
