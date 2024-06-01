package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Class struct {
	ID        primitive.ObjectID              `bson:"_id,omitempty" json:"id,omitempty"`
	Year      int                             `bson:"year" json:"year"`
	ClassName string                          `bson:"className" json:"className"`
	Divisions map[string][]primitive.ObjectID `bson:"divisions" json:"divisions"`
}

type ClassInfo struct {
	ClassName string          `json:"className"`
	Divisions map[string]bool `json:"divisions"`
}

type PostClassParams struct {
	Year      int                             `json:"year"`
	ClassName string                          `json:"className"`
	Divisions map[string][]primitive.ObjectID `json:"divisions"`
}

func (params PostClassParams) Validate() map[string]string {
	errors := map[string]string{}

	if params.Year == 0 {
		errors["year"] = "year is required"
	}
	if len(params.ClassName) < minClassName {
		errors["className"] = fmt.Sprintf("className length should be at least %d characters", minClassName)
	}
	if len(params.Divisions) == 0 {
		errors["divisions"] = "divisions are required"
	}
	return errors
}

func NewClassFromParams(params PostClassParams) (*Class, error) {
	return &Class{
		Year:      params.Year,
		ClassName: params.ClassName,
		Divisions: params.Divisions,
	}, nil
}
