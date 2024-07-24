package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	apierror "github.com/mayank12gt/fealtyx_assignment/error"
	"github.com/mayank12gt/fealtyx_assignment/repository"
	"github.com/mayank12gt/fealtyx_assignment/validator"
)

type StudentService struct {
	StudentRepo *repository.StudentRepo
}

type Query struct {
	Name     string
	AgeRange AgeRange
	Email    string
	PageSize int
	Page     int
}

type AgeRange struct {
	AgeMin int
	AgeMax int
}

func (s *StudentService) CreateStudent(student *repository.Student) *apierror.APIError {

	//validate request body
	if err := validator.ValidateStudent(student); err != nil {
		return err
	}

	return s.StudentRepo.CreateStudent(student)

}

func (s *StudentService) GetStudents(query Query) (*repository.Response, *apierror.APIError) {

	//validate Query Params
	if err := validateQueryParams(query); err != nil {
		return nil, err
	}

	students, meta, err := s.StudentRepo.GetStudents(query.Name, query.Email, query.AgeRange.AgeMin, query.AgeRange.AgeMax, query.Page, query.PageSize)
	if err != nil {
		return nil, err
	}

	response := repository.Response{
		Students: students,
		Metadata: *meta,
	}

	return &response, nil

}

func (s *StudentService) GetStudent(id int) (*repository.Student, *apierror.APIError) {

	return s.StudentRepo.GetStudent(id)

}

func (s *StudentService) DeleteStudent(id int) *apierror.APIError {

	return s.StudentRepo.DeleteStudent(id)

}

func (s *StudentService) UpdateStudent(id int, st *repository.Student) (*repository.Student, *apierror.APIError) {

	//get original student from repository
	student, err := s.StudentRepo.GetStudent(id)
	if err != nil {
		return nil, err
	}

	//update original student object's values with new values accordingly
	//after this we get student object with the updated values
	if st.Name != "" && len(st.Name) != 0 {
		student.Name = st.Name
	}
	if st.Email != "" && len(st.Email) != 0 {
		student.Email = st.Email
	}
	if st.Age != 0 {
		student.Age = st.Age
	}

	//validate the updated student object with new values
	//if the updated values are not valid, error will be returned and updates will not be saved
	if err := validator.ValidateStudent(student); err != nil {
		return nil, err
	}

	//call repository to save the updates
	updatedStudent, err := s.StudentRepo.UpdateStudent(id, student)
	if err != nil {
		return nil, err
	}

	return updatedStudent, nil

}

func (s *StudentService) GenerateSummary(student *repository.Student, baseUrl string) (*string, *apierror.APIError) {
	//Ollama Request body struct
	type LlamaReq struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
	}

	//creating prompt and request body
	prompt := fmt.Sprintf("Write a short summary for a student with name=%s age=%d email=%s in 20 words or less", student.Name, student.Age, student.Email)

	req := LlamaReq{
		Model:  "llama3",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {

		return nil, apierror.NewAPIError(500, "Error generating summary")
	}

	if baseUrl == "" || len(baseUrl) == 0 {
		return nil, apierror.NewAPIError(400, `Ollama Base URL is not set, restart the api with the Ollama_URL flag`)
	}

	//calling the Ollama api with the request body
	res, err := http.Post(baseUrl+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, apierror.NewAPIError(500, err.Error())
	}

	defer res.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, apierror.NewAPIError(500, err.Error())
	}

	//get the summary from response
	summary := string(body)

	return &summary, nil

}

func validateQueryParams(query Query) *apierror.APIError {

	if query.Page < 1 {
		return apierror.NewAPIError(422, "Page must be >= 1")
	}

	if !validator.ValidateIntegerRange(query.PageSize, 1, 20) {
		return apierror.NewAPIError(422, "Page Size must be between 1 and 20")
	}

	if validator.ValidateIntegerRange(query.AgeRange.AgeMin, 1, 100) {
		return apierror.NewAPIError(422, "ageMin must be between 1 and 100")
	}

	if validator.ValidateIntegerRange(query.AgeRange.AgeMax, 1, 100) {
		return apierror.NewAPIError(422, "ageMax must be between 1 and 100")
	}

	if query.AgeRange.AgeMin > query.AgeRange.AgeMax {
		return apierror.NewAPIError(422, "ageMin must be less than ageMax")
	}

	return nil
}
