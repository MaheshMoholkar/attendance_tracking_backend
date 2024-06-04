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

type ClassDivision struct {
	ClassName string   `json:"className"`
	Divisions []string `json:"divisions"`
}

type ClassDivisions struct {
	ClassNames []string
	Divisions  map[string][]string
}

func ParseClassDivisions(dbClassDivisions []postgres.GetClassDivisionsRow) ClassDivisions {
	var classDivisions ClassDivisions
	classDivisionMap := make(map[string][]string)
	var classNames []string

	for _, cd := range dbClassDivisions {
		// Add classNames to the list
		if _, exists := classDivisionMap[cd.Classname]; !exists {
			classNames = append(classNames, cd.Classname)
		}
		classDivisionMap[cd.Classname] = append(classDivisionMap[cd.Classname], cd.Divisionname)
	}

	classDivisions = ClassDivisions{
		ClassNames: classNames,
		Divisions:  classDivisionMap,
	}
	return classDivisions
}
