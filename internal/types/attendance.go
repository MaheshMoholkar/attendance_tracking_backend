package types

type AttendanceUpdate struct {
	StudentID  int               `json:"student_id"`
	Attendance map[string]string `json:"attendance"`
}
