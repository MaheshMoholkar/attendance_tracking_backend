package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Class struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ClassName string             `bson:"className" json:"className"`
}

type PostClassParams struct {
	ClassName string `json:"className"`
}

func (params PostClassParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.ClassName) < minClassName {
		errors["className"] = fmt.Sprintf("className length should be at least %d characters", minClassName)
	}
	return errors
}
