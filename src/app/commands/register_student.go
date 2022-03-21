package commands

import (
	"encoding/json"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type RegisterStudentCommand struct {
	db *gorm.DB
	nc *nats.Conn
}

func NewRegisterStudentCommand(db *gorm.DB, nc *nats.Conn) RegisterStudentCommand {
	return RegisterStudentCommand{
		db: db,
		nc: nc,
	}
}

func (rs *RegisterStudentCommand) Execute(student *entities.Student) error {
	if err := student.Prepare(); err != nil {
		return err
	}

	if err := rs.db.Create(student).Error; err != nil {
		return err
	}

	jason, err := json.Marshal(student)

	if err != nil {
		return err
	}

	rs.nc.Publish("register_student", jason)

	return nil
}
