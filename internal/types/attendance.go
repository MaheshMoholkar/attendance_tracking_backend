package types

import (
	"time"

	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
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
	StudentID  int32                  `json:"student_id"`
	Attendance map[string]interface{} `json:"attendance"`
}

// Define the data structure for the response
type StudentAttendance struct {
	StudentID  int          `json:"student_id"`
	Attendance map[int]bool `json:"attendance"`
}

// Assuming postgres.AttendanceInfo has the same fields as types.Attendance
type Attendance struct {
	AttendanceID int
	StudentID    int
	Date         time.Time
	Status       bool
	ClassID      int
	DivisionID   int
}

func ParseAttendance(dbAttendance postgres.AttendanceInfo) Attendance {
	return Attendance{
		AttendanceID: int(dbAttendance.AttendanceID),
		StudentID:    int(dbAttendance.StudentID),
		Date:         dbAttendance.Date,
		Status:       dbAttendance.Status,
		ClassID:      int(dbAttendance.ClassID),
		DivisionID:   int(dbAttendance.DivisionID),
	}
}

func ParseAttendances(dbAttendances []postgres.AttendanceInfo) []Attendance {
	attendances := []Attendance{}
	for _, dbAttendance := range dbAttendances {
		attendances = append(attendances, ParseAttendance(dbAttendance))
	}
	return attendances
}

func ConvertAttendanceData(attendances []Attendance) []StudentAttendance {
	attendanceMap := make(map[int]map[int]bool)
	for _, a := range attendances {
		day := a.Date.Day()
		if _, exists := attendanceMap[a.StudentID]; !exists {
			attendanceMap[a.StudentID] = make(map[int]bool)
		}
		attendanceMap[a.StudentID][day] = a.Status
	}

	result := []StudentAttendance{}
	for studentID, days := range attendanceMap {
		result = append(result, StudentAttendance{
			StudentID:  studentID,
			Attendance: days,
		})
	}

	return result
}
