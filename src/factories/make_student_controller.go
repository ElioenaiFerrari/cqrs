package factories

import (
	"github.com/ElioenaiFerrari/cqrs/src/app/commands"
	"github.com/ElioenaiFerrari/cqrs/src/presentation/controllers"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
)

func MakeStudentController(db *gorm.DB, producer *kafka.Producer) controllers.StudentsController {
	registerStudentCommand := commands.NewRegisterStudentCommand(db, producer)

	return controllers.NewStudentsController(registerStudentCommand)

}
