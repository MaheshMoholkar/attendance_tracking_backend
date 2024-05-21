package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Class struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ClassName string             `bson:"className" json:"className"`
	Divisions []string           `bson:"divisions" json:"divisions"`
}

type PostClassParams struct {
	ClassName string   `json:"className"`
	Divisions []string `json:"divisions"`
}

func (params PostClassParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.ClassName) < minClassName {
		errors["className"] = fmt.Sprintf("className length should be at least %d characters", minClassName)
	}
	if len(params.Divisions) == 0 {
		errors["divisions"] = "divisions are required"
	}
	for _, division := range params.Divisions {
		if len(division) != 1 {
			errors["division"] = "division must be one character long"
		}
	}
	return errors
}

func NewClassFromParams(params PostClassParams) (*Class, error) {
	return &Class{
		ClassName: params.ClassName,
		Divisions: params.Divisions,
	}, nil
}
