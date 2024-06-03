package types

import "github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"

type ClassInfo struct {
	ClassName string `json:"className"`
}

func NewClass(params ClassInfo) *ClassInfo {
	return &ClassInfo{
		ClassName: params.ClassName,
	}
}

func ParseClass(dbClass postgres.ClassInfo) ClassInfo {
	return ClassInfo{
		ClassName: dbClass.Classname,
	}
}

func ParseClasses(dbClasses []postgres.ClassInfo) []ClassInfo {
	classes := make([]ClassInfo, len(dbClasses))
	for i, dbClass := range dbClasses {
		classes[i] = ParseClass(dbClass)
	}
	return classes
}
