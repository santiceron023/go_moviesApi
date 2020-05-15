package validator

import (
	"errors"
	"strings"
)

func ValidateRequired(value string, message string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New(message)
	}
	return nil
}

func ValidateLength(value int,message string) error{
	if value <= 0 {
		return errors.New(message)
	}
	return nil
}