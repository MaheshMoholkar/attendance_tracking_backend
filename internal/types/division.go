package types

import "github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"

type Division struct {
	DivisionID   int32  `json:"division_id"`
	DivisionName string `json:"division_name"`
	ClassID      int32  `json:"class_id"`
}

func NewDivision(params Division) *Division {
	return &Division{
		DivisionID:   params.DivisionID,
		DivisionName: params.DivisionName,
		ClassID:      params.ClassID,
	}
}

func ParseDivision(dbDivision postgres.DivisionInfo) Division {
	return Division{
		DivisionID:   dbDivision.DivisionID,
		DivisionName: dbDivision.DivisionName,
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
