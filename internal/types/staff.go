package types

import (
	"github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"
)

type Staff struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	StaffID   int32  `json:"staff_id"`
	Password  string `json:"password"`
}

func NewStaff(params Staff) *Staff {
	return &Staff{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		StaffID:   params.StaffID,
	}
}

func ParseStaff(dbStaff postgres.Staff) Staff {
	return Staff{
		FirstName: dbStaff.Firstname,
		LastName:  dbStaff.Lastname,
		Email:     dbStaff.Email,
		StaffID:   dbStaff.StaffID,
	}
}

func ParseStaffs(dbStaffs []postgres.Staff) []Staff {
	staffs := make([]Staff, len(dbStaffs))
	for i, dbStaff := range dbStaffs {
		staffs[i] = ParseStaff(dbStaff)
	}
	return staffs
}
