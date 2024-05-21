package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Rollno    int                `bson:"rollno" json:"rollno"`
	Email     string             `bson:"email" json:"email"`
	ClassName string             `bson:"className" json:"className"`
	Division  string             `bson:"division" json:"division"`
	Year      int                `bson:"year" json:"year"`
}

type PostStudentParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Rollno    int    `json:"rollno"`
	Email     string `json:"email"`
	ClassName string `json:"className"`
	Division  string `json:"division"`
	Year      int    `json:"year"`
}

func (params PostStudentParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstName {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstName)
	}
	if len(params.LastName) < minLastName {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLastName)
	}
	if params.Rollno == 0 {
		errors["rollno"] = "rollno is required"
	}
	if !isEmailValid(params.Email) {
		errors["email"] = "email is invalid"
	}
	if len(params.ClassName) < minClassName {
		errors["className"] = "className is required"
	}
	if len(params.Division) == 0 {
		errors["division"] = "division is required"
	}
	if params.Year == 0 {
		errors["year"] = "year is required"
	}
	return errors
}

func NewStudentFromParams(params PostStudentParams) (*Student, error) {
	return &Student{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Rollno:    params.Rollno,
		Email:     params.Email,
		ClassName: params.ClassName,
		Division:  params.Division,
		Year:      params.Year,
	}, nil
}
