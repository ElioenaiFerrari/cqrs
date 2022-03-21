package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/cqrs/src/app/commands"
	"github.com/ElioenaiFerrari/cqrs/src/app/queries"
	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
	"github.com/labstack/echo/v4"
)

type StudentsController struct {
	registerStudentCommand commands.RegisterStudentCommand
	getStudentsQuery       queries.GetStudentsQuery
}

func NewStudentsController(registerStudentCommand commands.RegisterStudentCommand, getStudentsQuery queries.GetStudentsQuery) StudentsController {
	return StudentsController{
		registerStudentCommand: registerStudentCommand,
		getStudentsQuery:       getStudentsQuery,
	}
}

func (sc *StudentsController) Create(c echo.Context) error {
	var student entities.Student

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if err := sc.registerStudentCommand.Execute(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}

func (sc *StudentsController) Index(c echo.Context) error {
	var students []entities.Student

	students, err := sc.getStudentsQuery.Execute()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, students)
}
