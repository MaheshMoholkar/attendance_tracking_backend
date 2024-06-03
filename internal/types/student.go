package types

import (
	"fmt"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
)

type Student struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Rollno    int32  `json:"rollno"`
	Email     string `json:"email"`
	ClassName string `json:"className"`
	Division  string `json:"division"`
	Year      int32  `json:"year"`
	StudentID int32  `json:"student_id"`
}

func (params Student) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstName {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstName)
	}
	if len(params.LastName) < minLastName {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLastName)
	}
	if params.Rollno == 0 {
		errors["rollno"] = "rollno is required"
	}
	if !isEmailValid(params.Email) {
		errors["email"] = "email is invalid"
	}
	if len(params.ClassName) < minClassName {
		errors["className"] = "className is required"
	}
	if len(params.Division) == 0 {
		errors["division"] = "division is required"
	}
	if params.Year == 0 {
		errors["year"] = "year is required"
	}
	if params.Year == 0 {
		errors["student_id"] = "studentID is required"
	}
	return errors
}

func NewStudent(params Student) *Student {
	return &Student{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Rollno:    params.Rollno,
		Email:     params.Email,
		ClassName: params.ClassName,
		Division:  params.Division,
		Year:      params.Year,
		StudentID: params.StudentID,
	}
}

func ParseStudent(dbStudent postgres.StudentInfo) Student {
	return Student{
		FirstName: dbStudent.Firstname,
		LastName:  dbStudent.Lastname,
		Rollno:    dbStudent.Rollno,
		Email:     dbStudent.Email,
		ClassName: dbStudent.Classname,
		Division:  dbStudent.Division,
		Year:      dbStudent.Year,
		StudentID: dbStudent.StudentID,
	}
}

func ParseStudents(dbStudents []postgres.StudentInfo) []Student {
	students := []Student{}
	for _, dbStudent := range dbStudents {
		students = append(students, ParseStudent(dbStudent))
	}
	return students
}
