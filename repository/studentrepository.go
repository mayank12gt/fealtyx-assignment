package repository

import (
	apierror "github.com/mayank12gt/fealtyx_assignment/error"
)

type StudentRepo struct {
	// generally we'll have a db connection here

	Students []Student
}

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (r *StudentRepo) CreateStudent(student *Student) *apierror.APIError {

	// we'll make db call here and handle any returns accordingly
	if len(r.Students) > 0 {
		student.ID = r.Students[len(r.Students)-1].ID + 1
	} else {
		student.ID = 1
	}

	r.Students = append(r.Students, *student)

	return nil
}

func (r *StudentRepo) GetStudents() ([]Student, *apierror.APIError) {

	// we'll make db call here and handle any returns accordingly

	return r.Students, nil

}

func (r *StudentRepo) GetStudent(id int) (*Student, *apierror.APIError) {

	// we'll make db call here and handle any returns accordingly

	for _, student := range r.Students {
		if student.ID == id {
			return &student, nil
		}
	}

	return nil, apierror.NewAPIError(404, "Student not found")

}

func (r *StudentRepo) DeleteStudent(id int) *apierror.APIError {

	// we'll make db call here and handle any returns accordingly

	for idx, student := range r.Students {
		if student.ID == id {
			r.Students = append(r.Students[:idx], r.Students[idx+1:]...)
			return nil
		}
	}

	return apierror.NewAPIError(404, "Student not found")

}

func (r *StudentRepo) UpdateStudent(id int, student *Student) (*Student, *apierror.APIError) {

	// we'll make db call here and handle any returns accordingly

	for idx, st := range r.Students {
		if st.ID == id {

			r.Students[idx] = *student

			return &r.Students[idx], nil
		}
	}

	return nil, apierror.NewAPIError(404, "Student not found")

}
