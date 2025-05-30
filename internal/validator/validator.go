package validator

import (
	"errors"
	"strings"
)

type CreateUserInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int32  `json:"age"`
}

func (i *CreateUserInput) Validate() error {
	if strings.TrimSpace(i.FullName) == "" || len(i.FullName) > 255 {
		return errors.New("invalid full name")
	}
	if strings.TrimSpace(i.Email) == "" || len(i.Email) > 255 {
		return errors.New("invalid email")
	}
	if i.Age < 18 {
		return errors.New("age must be >= 18")
	}
	return nil
}

type UpdateUserInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int32  `json:"age"`
}

func (i *UpdateUserInput) Validate() error {
	return (&CreateUserInput{
		FullName: i.FullName,
		Email:    i.Email,
		Age:      i.Age,
	}).Validate()
}
