package tools

import (
	"fmt"
)

// FieldError is a struct that represents a field error
type FieldError struct {
	// Field is the field that has the error
	Field string
	// Msg is the error message
	Msg string
}

// Error is a method that returns the error message
func (e *FieldError) Error() string {
	return fmt.Sprintf("field %s: %s", e.Field, e.Msg)
}

// ValidateField is a function that validates the required
func ValidateField(fields map[string]any, required ...string) (err error) {
	for _, field := range required {
		if _, ok := fields[field]; !ok {
			err = &FieldError{Field: field, Msg: "is required"}
			return
		}
	}
	return
}
