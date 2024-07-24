package repository

import (
	"log"
	"math"
	"strings"

	apierror "github.com/mayank12gt/fealtyx_assignment/error"
)

type StudentRepo struct {
	// generally we'll have a db connection here

	Students []Student
}

type Response struct {
	Metadata Metadata  `json:"metadata"`
	Students []Student `json:"students"`
}

type Metadata struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	FirstPage    int `json:"first_page"`
	LastPage     int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

func (r *StudentRepo) CreateStudent(student *Student) *apierror.APIError {

	// we'll make db call here and handle any returns accordingly
	if len(r.Students) > 0 {
		student.ID = r.Students[len(r.Students)-1].ID + 1
	} else {
		student.ID = 1
	}

	for _, st := range r.Students {
		if st.Email == student.Email {
			return apierror.NewAPIError(400, "Email already exists")
		}
	}

	r.Students = append(r.Students, *student)

	return nil
}

func (r *StudentRepo) GetStudents(name, email string, ageMin, ageMax, page, pageSize int) ([]Student, *Metadata, *apierror.APIError) {

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	// we'll make db call here and handle any returns accordingly
	var students []Student
	count := 0
	for _, student := range r.Students {
		log.Print(student)
		// Apply name filter if name is not empty
		if name != "" && !strings.Contains(strings.ToLower(student.Name), strings.ToLower(name)) {
			continue
		}

		// Apply email filter if email is not empty
		if email != "" && student.Email != email {
			continue
		}

		// Apply age range filter
		if student.Age < ageMin || student.Age > ageMax {
			continue
		}

		students = append(students, student)
		count++
	}

	if count == 0 {
		return nil, nil, apierror.NewAPIError(404, "Students not found")
	}

	meta := calculateMetadata(count, int(page), int(pageSize))
	// log.Print(meta)

	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex > len(r.Students)-1 {
		return []Student{}, &meta, nil
	}

	// log.Print(startIndex)
	// log.Print(endIndex)

	if startIndex >= len(students) {
		students = []Student{}
		// Return an empty slice if the start index is out of bounds
	} else if endIndex > len(students) {
		endIndex = len(students)
		// Adjust the end index if it exceeds the slice length
	}

	students = students[startIndex:endIndex]

	// log.Print(students)

	return students, &meta, nil

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
