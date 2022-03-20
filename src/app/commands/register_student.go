package commands

import (
	"encoding/json"

	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
)

type RegisterStudentCommand struct {
	producer *kafka.Producer
	db       *gorm.DB
}

func NewRegisterStudentCommand(db *gorm.DB, producer *kafka.Producer) RegisterStudentCommand {
	return RegisterStudentCommand{
		db:       db,
		producer: producer,
	}
}

func (rs *RegisterStudentCommand) Execute(student *entities.Student) error {
	if err := student.Prepare(); err != nil {
		return err
	}

	if err := rs.db.Create(student).Error; err != nil {
		return err
	}

	defer rs.producer.Close()

	jason, err := json.Marshal(student)

	if err != nil {
		return err
	}

	topic := "register_student"

	rs.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:          jason,
	}, nil)

	rs.producer.Flush(10 * 1000)

	return nil
}
