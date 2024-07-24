package validator

import (
	apierror "github.com/mayank12gt/fealtyx_assignment/error"
	"github.com/mayank12gt/fealtyx_assignment/repository"
)

func ValidateStudent(student *repository.Student) *apierror.APIError {

	if !ValidateIntegerRange(len(student.Name), 3, 50) {
		return apierror.NewAPIError(422, "Name must be between 3 and 50 characters")
	}

	if !ValidateIntegerRange(student.Age, 1, 100) {
		return apierror.NewAPIError(422, "Age must be between 1 and 100")
	}

	if !ValidateEmail(student.Email) {
		return apierror.NewAPIError(422, "Invalid email id")
	}

	return nil
}
