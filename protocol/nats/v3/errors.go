package nats

import "fmt"

// validationError is returned when an invalid option is given
type validationError struct {
	field   string
	message string
}

// Error returns a message indicating an error condition, with the nil value representing no error.
func (v validationError) Error() string {
	return fmt.Sprintf("invalid parameters provided: %q: %s", v.field, v.message)
}

// newValidationError creates a validation error
func newValidationError(field, message string) validationError {
	return validationError{
		field:   field,
		message: message,
	}
}
