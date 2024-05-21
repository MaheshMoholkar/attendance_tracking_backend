package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendance struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Rollno   int                `bson:"rollno" json:"rollno"`
	Present  bool               `bson:"present" json:"present"`
	DateTime primitive.DateTime `bson:"dateTime" json:"dateTime"`
}

type PostAttendanceParams struct {
	Rollno   int                `json:"rollno"`
	Present  bool               `json:"present"`
	DateTime primitive.DateTime `json:"dateTime"`
}

func (params PostAttendanceParams) Validate() map[string]string {
	errors := map[string]string{}
	if params.Rollno == 0 {
		errors["rollo"] = "rollno is required"
	}
	dateTime := params.DateTime.Time()
	if dateTime.IsZero() {
		errors["dateTime"] = "dateTime is required"
	} else if !isValidDateTime(dateTime) {
		errors["dateTime"] = "dateTime must be a valid date and time"
	}

	return errors
}

func isValidDateTime(t time.Time) bool {
	return !t.IsZero()
}

func NewAttendaceFromParams(params PostAttendanceParams) (*Attendance, error) {
	return &Attendance{
		Rollno:   params.Rollno,
		Present:  params.Present,
		DateTime: params.DateTime,
	}, nil
}
