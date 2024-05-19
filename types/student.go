package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Class     string             `bson:"class" json:"class"`
	Division  string             `bson:"division" json:"division"`
}

type PostStudentParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Class     string `json:"class"`
	Division  string `json:"division"`
}

func (params PostStudentParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstName {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstName)
	}
	if len(params.LastName) < minLastName {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLastName)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = "email is invalid"
	}
	if len(params.Class) == 0 {
		errors["class"] = "class is needed"
	}
	if len(params.Division) == 0 {
		errors["division"] = "division is needed"
	}
	return errors
}

func NewStudentFromParams(params PostStudentParams) (*Student, error) {
	return &Student{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Class:     params.Class,
		Division:  params.Division,
	}, nil
}
