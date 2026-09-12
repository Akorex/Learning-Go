package main

import (
	"errors"
	"fmt"
)

type ValidationPerson struct {
	FirstName string
	LastName  string
	Age       int
}

func ValidatePerson(p ValidationPerson) error {
	var errs []error

	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("field FirstName cannot be empty"))
	}

	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("field LastName cannot be empty"))
	}

	if p.Age < 0 {
		errs = append(errs, errors.New("field Age cannot be negative"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func Run21ErrorJoiningValidation() {
	p := ValidationPerson{
		FirstName: "",
		LastName:  "",
		Age:       -5,
	}

	err := ValidatePerson(p)
	if err != nil {
		fmt.Println("Validation errors:")
		fmt.Println(err)
	} else {
		fmt.Println("Person is valid")
	}
}
