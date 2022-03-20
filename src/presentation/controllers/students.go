package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/cqrs/src/app/commands"
	"github.com/ElioenaiFerrari/cqrs/src/domain/entities"
	"github.com/labstack/echo/v4"
)

type StudentsController struct {
	registerStudentCommand commands.RegisterStudentCommand
}

func NewStudentsController(registerStudentCommand commands.RegisterStudentCommand) StudentsController {
	return StudentsController{
		registerStudentCommand: registerStudentCommand,
	}
}

func (sc *StudentsController) Create(c echo.Context) error {
	var student entities.Student

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := sc.registerStudentCommand.Execute(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, student)
}
