package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

type PostUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params PostUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstName {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstName)
	}
	if len(params.LastName) < minLastName {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLastName)
	}
	if len(params.Password) < minPassword {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPassword)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = "email is invalid"
	}
	return errors
}

func NewUserFromParams(params PostUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
