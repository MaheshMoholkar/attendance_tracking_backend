package types

import "github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"

type Division struct {
	DivisionName string `json:"division_name"`
	ClassID      int32  `json:"class_id"`
}

func NewDivision(params Division) *Division {
	return &Division{
		DivisionName: params.DivisionName,
		ClassID:      params.ClassID,
	}
}

func ParseDivision(dbDivision postgres.DivisionInfo) Division {
	return Division{
		DivisionName: dbDivision.Divisionname,
		ClassID:      dbDivision.ClassID,
	}
}

func ParseDivisions(dbDivisions []postgres.DivisionInfo) []Division {
	divisions := make([]Division, len(dbDivisions))
	for i, dbDivision := range dbDivisions {
		divisions[i] = ParseDivision(dbDivision)
	}
	return divisions
}
