package main

import (
	"errors"
	"fmt"
)

func userIsEligible(email, password string, age int) error {
	if email == "" {
		return errors.New("email can't be empty")
	}
	if password == "" {
		return errors.New("password can't be empty")
	}
	const minAge = 18
	if age < 18 {
		return fmt.Errorf("age must be at least %v years old", minAge)
	}
	return nil
}
