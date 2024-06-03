package types

import (
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

func ParseStudent(dbStudent postgres.Student) Student {
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

func ParseStudents(dbStudents []postgres.Student) []Student {
	students := []Student{}
	for _, dbStudent := range dbStudents {
		students = append(students, ParseStudent(dbStudent))
	}
	return students
}
