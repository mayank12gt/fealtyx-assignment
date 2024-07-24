package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
	apierror "github.com/mayank12gt/fealtyx_assignment/error"
	"github.com/mayank12gt/fealtyx_assignment/repository"
)

func (app *App) HealthCheckHandler() func(c echo.Context) error {
	return func(c echo.Context) error {

		return c.JSON(200, map[string]string{
			"message": "up and running",
		})
	}

}

func (app *App) GetStudentsHandler() func(c echo.Context) error {
	return func(c echo.Context) error {

		students, err := app.StudentService.GetStudents()
		if err != nil {
			return c.JSON(err.Code, err)
		}

		return c.JSON(200, students)
	}

}

func (app *App) CreateStudentHandler() func(c echo.Context) error {
	return func(c echo.Context) error {

		var student repository.Student

		if err := c.Bind(&student); err != nil {
			apiErr := apierror.NewAPIError(400, "Bad Request. Verify Request Body")
			return c.JSON(apiErr.Code, apiErr)
		}

		//call student service and add student to db

		if err := app.StudentService.CreateStudent(&student); err != nil {
			return c.JSON(err.Code, err)
		}

		return c.JSON(200, map[string]string{
			"message": "Student Created Successfully",
		})
	}

}

func (app *App) GetStudentHandler() func(c echo.Context) error {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			apiErr := apierror.NewAPIError(400, "Id must be integer")

			return c.JSON(apiErr.Code, apiErr)
		}

		//validate ID
		if id < 1 {
			apiErr := apierror.NewAPIError(400, "Id must be > 0")

			return c.JSON(apiErr.Code, apiErr)
		}

		//call repository
		student, apiErr := app.StudentService.GetStudent(id)
		if apiErr != nil {
			return c.JSON(apiErr.Code, apiErr)
		}

		return c.JSON(200, student)
	}

}

func (app *App) DeleteStudentHandler() func(c echo.Context) error {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			apiErr := apierror.NewAPIError(400, "Id must be integer")

			return c.JSON(apiErr.Code, apiErr)
		}

		//validate ID
		if id < 1 {
			apiErr := apierror.NewAPIError(400, "Id must be > 0")

			return c.JSON(apiErr.Code, apiErr)
		}

		//call repository
		apiErr := app.StudentService.DeleteStudent(id)
		if apiErr != nil {
			return c.JSON(apiErr.Code, apiErr)
		}

		return c.JSON(200, map[string]string{
			"message": "Student deleted successfully",
		})
	}

}

func (app *App) UpdateStudentHandler() func(c echo.Context) error {
	return func(c echo.Context) error {

		var student repository.Student

		if err := c.Bind(&student); err != nil {
			apiErr := apierror.NewAPIError(400, "Bad Request. Verify Request Body")
			return c.JSON(apiErr.Code, apiErr)
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			apiErr := apierror.NewAPIError(400, "Id must be integer")

			return c.JSON(apiErr.Code, apiErr)
		}

		//validate id
		if id < 1 {
			apiErr := apierror.NewAPIError(400, "Id must be > 0")

			return c.JSON(apiErr.Code, apiErr)
		}

		//call student service and add student to db

		updatedStudent, apiErr := app.StudentService.UpdateStudent(id, &student)
		if apiErr != nil {
			return c.JSON(apiErr.Code, apiErr)
		}

		return c.JSON(200, updatedStudent)

	}
}

func (app *App) GetStudentSummaryHandler() func(c echo.Context) error {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			apiErr := apierror.NewAPIError(400, "Id must be integer")

			return c.JSON(apiErr.Code, apiErr)
		}

		//validate id
		if id < 1 {
			apiErr := apierror.NewAPIError(422, "Id must be > 0")

			return c.JSON(apiErr.Code, apiErr)
		}

		//call service to get student
		student, apiErr := app.StudentService.GetStudent(id)
		if apiErr != nil {

			return c.JSON(apiErr.Code, apiErr)
		}

		//call generate summary service
		summary, apiErr := app.StudentService.GenerateSummary(student, app.Config.Ollama_Base_Url)
		if apiErr != nil {

			return c.JSON(apiErr.Code, apiErr)
		}

		return c.JSON(200, map[string]string{
			"summary": *summary,
		})

	}
}
