package types

type AttendanceUpdate struct {
	RowData      []StudentAttendance `json:"rowData"`
	ClassName    string              `json:"className"`
	DivisionName string              `json:"divisionName"`
	MonthYear    string              `json:"monthYear"`
	Subject      string              `json:"subject"`
}

type StudentAttendance struct {
	StudentID  int          `json:"student_id"`
	Attendance map[int]bool `json:"attendance"`
}
