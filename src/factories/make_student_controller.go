package factories

import (
	"github.com/ElioenaiFerrari/cqrs/src/app/commands"
	"github.com/ElioenaiFerrari/cqrs/src/presentation/controllers"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func MakeStudentController(db *gorm.DB, nc *nats.Conn) controllers.StudentsController {
	registerStudentCommand := commands.NewRegisterStudentCommand(db, nc)

	return controllers.NewStudentsController(registerStudentCommand)

}
