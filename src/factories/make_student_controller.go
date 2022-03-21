package factories

import (
	"github.com/ElioenaiFerrari/cqrs/src/app/commands"
	"github.com/ElioenaiFerrari/cqrs/src/app/queries"
	"github.com/ElioenaiFerrari/cqrs/src/presentation/controllers"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func MakeStudentController(writeDB *gorm.DB, readDB *mongo.Database, nc *nats.Conn) controllers.StudentsController {
	registerStudentCommand := commands.NewRegisterStudentCommand(writeDB, nc)
	getStudentsQuery := queries.NewGetStudentsQuery(readDB)

	return controllers.NewStudentsController(registerStudentCommand, getStudentsQuery)

}
