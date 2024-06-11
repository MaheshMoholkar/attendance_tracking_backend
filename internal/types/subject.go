package types

import "github.com/MaheshMoholkar/attendance_tracking_backend/internal/database/postgres"

type Subject struct {
	SubjectName string `json:"subject_name"`
	ClassID     int32  `json:"class_id"`
}

func NewSubject(params Subject) *Subject {
	return &Subject{
		SubjectName: params.SubjectName,
		ClassID:     params.ClassID,
	}
}

func ParseSubject(dbSubject postgres.SubjectInfo) Subject {
	return Subject{
		SubjectName: dbSubject.Subjectname,
		ClassID:     dbSubject.ClassID,
	}
}

func ParseSubjects(dbSubjects []postgres.SubjectInfo) []Subject {
	subjects := make([]Subject, len(dbSubjects))
	for i, dbSubject := range dbSubjects {
		subjects[i] = ParseSubject(dbSubject)
	}
	return subjects
}
