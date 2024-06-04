package types

import (
	"fmt"
	"time"
)

// AttendanceList represents the payload for saving attendance.
type AttendanceList struct {
	Data  []AttendanceMap `json:"data"`
	Month int             `json:"month"`
	Year  int             `json:"year"`
}

// AttendanceMap represents the dynamic structure of attendance data.
type AttendanceMap struct {
	Name       string                 `json:"name"`
	StudentID  int32                  `json:"studentID"`
	ClassID    int32                  `json:"classID"`
	DivisionID int32                  `json:"divisionID"`
	Attendance map[string]interface{} `json:"attendance"`
}

// Attendance represents a record in the attendance_info table.
type Attendance struct {
	AttendanceID int32     `json:"attendanceID"`
	StudentID    int32     `json:"studentID"`
	Date         time.Time `json:"date"`
	Status       bool      `json:"status"`
	ClassID      int32     `json:"classID"`
	DivisionID   int32     `json:"divisionID"`
}

// Validate checks the fields of Attendance for validity.
func (params Attendance) Validate() map[string]string {
	errors := map[string]string{}

	// Check if StudentID is provided
	if params.StudentID == 0 {
		errors["studentID"] = "studentID is required"
	}

	// Check if Date is valid
	if params.Date.IsZero() {
		errors["date"] = "date is required"
	} else if !isValidDate(params.Date) {
		errors["date"] = "date must be a valid date"
	}

	// Check if ClassID is provided
	if params.ClassID == 0 {
		errors["classID"] = "classID is required"
	}

	// Check if DivisionID is provided
	if params.DivisionID == 0 {
		errors["divisionID"] = "divisionID is required"
	}

	return errors
}

// isValidDate checks if the provided time is a valid date.
func isValidDate(t time.Time) bool {
	return !t.IsZero()
}

// PostAttendanceParams represents the parameters for creating an Attendance record.
type PostAttendanceParams struct {
	StudentID  int32     `json:"studentID"`
	Date       time.Time `json:"date"`
	Status     bool      `json:"status"`
	ClassID    int32     `json:"classID"`
	DivisionID int32     `json:"divisionID"`
}

// NewAttendanceFromParams creates a new Attendance instance from the provided parameters.
func NewAttendanceFromParams(params PostAttendanceParams) (*Attendance, error) {
	attendance := &Attendance{
		StudentID:  params.StudentID,
		Date:       params.Date,
		Status:     params.Status,
		ClassID:    params.ClassID,
		DivisionID: params.DivisionID,
	}

	// Validate the attendance parameters
	errors := attendance.Validate()
	if len(errors) > 0 {
		return nil, fmt.Errorf("validation errors: %v", errors)
	}

	return attendance, nil
}
