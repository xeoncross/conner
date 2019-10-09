package conner

import (
	"errors"
)

// Map of fields to append to the structured error when logging
// Alias to save typing
type Map map[string]interface{}

type structuredError struct {
	err    error
	fields map[string]interface{}
}

// Error message
func (err structuredError) Error() string {
	return err.err.Error()
}

// Unwrap error cause
func (err structuredError) Unwrap() error {
	return err.err
}

// Error with values
func Error(err error, fields map[string]interface{}) error {
	return structuredError{err: err, fields: fields}
}

// Unclear half-mix of fmt directives (required %w, but no others allowed)
// Can lead to runtime errors
// Errorf wrap error message and set error fields
// func Errorf(msg string, err error, fields map[string]interface{}) error {
// 	return structuredError{err: fmt.Errorf(msg, err), fields: fields}
// }

// Values from each error on the error chain
func Values(err error) map[string]interface{} {
	output := map[string]interface{}{}

	var e structuredError

	for errors.As(err, &e) {
		for field, val := range e.fields {
			output[field] = val
		}
		err = e.Unwrap()
	}

	return output
}
