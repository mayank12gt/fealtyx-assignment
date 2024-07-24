package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
	apierror "github.com/mayank12gt/fealtyx_assignment/error"
	"github.com/mayank12gt/fealtyx_assignment/repository"
	"github.com/mayank12gt/fealtyx_assignment/service"
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

		name := c.QueryParam("name")
		email := c.QueryParam("email")
		var ageMin, ageMax, page, pageSize int
		var err error
		if c.QueryParam("ageMin") != "" {
			ageMin, err = strconv.Atoi(c.QueryParam("ageMin"))
			if err != nil {
				apiErr := apierror.NewAPIError(400, "ageMin must be an integer")
				return c.JSON(apiErr.Code, apiErr)
			}
		} else {
			ageMin = 1
		}
		if c.QueryParam("ageMax") != "" {
			ageMax, err = strconv.Atoi(c.QueryParam("ageMax"))
			if err != nil {
				apiErr := apierror.NewAPIError(400, "ageMax must be an integer")
				return c.JSON(apiErr.Code, apiErr)
			}
		} else {
			ageMax = 100
		}

		if c.QueryParam("page") != "" {
			page, err = strconv.Atoi(c.QueryParam("page"))
			if err != nil {
				apiErr := apierror.NewAPIError(400, "page must be an integer")
				return c.JSON(apiErr.Code, apiErr)
			}
		} else {
			page = 1
		}

		if c.QueryParam("page_size") != "" {
			pageSize, err = strconv.Atoi(c.QueryParam("page_size"))
			if err != nil {
				apiErr := apierror.NewAPIError(400, "page size must be an integer")
				return c.JSON(apiErr.Code, apiErr)
			}
		} else {
			pageSize = 20
		}

		query := service.Query{
			Name:  name,
			Email: email,
			AgeRange: service.AgeRange{
				AgeMin: ageMin,
				AgeMax: ageMax,
			},
			Page:     page,
			PageSize: pageSize,
		}

		response, err2 := app.StudentService.GetStudents(query)
		if err2 != nil {
			return c.JSON(err2.Code, err2)
		}

		return c.JSON(200, response)
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
