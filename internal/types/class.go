package types

import "github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"

type Class struct {
	ClassID   int32  `json:"class_id"`
	ClassName string `json:"className"`
}

func NewClass(params Class) *Class {
	return &Class{
		ClassID:   params.ClassID,
		ClassName: params.ClassName,
	}
}

func ParseClass(dbClass postgres.ClassInfo) Class {
	return Class{
		ClassID:   dbClass.ClassID,
		ClassName: dbClass.Classname,
	}
}

func ParseClasses(dbClasses []postgres.ClassInfo) []Class {
	classes := make([]Class, len(dbClasses))
	for i, dbClass := range dbClasses {
		classes[i] = ParseClass(dbClass)
	}
	return classes
}
