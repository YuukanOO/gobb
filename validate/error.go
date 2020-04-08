package validate

import "fmt"

const (
	// ErrCode in the DomainError returned by Guard.
	ErrCode = "validation_failed"
	// ErrMessage in the DomainError returned by Guard.
	ErrMessage = "some validation has failed"
)

type (
	// ValidationErrors maps field name to field errors and implements
	// the Error interface.
	ValidationErrors map[string]string

	// FieldError represents an error contextualized to a field.
	FieldError struct {
		Field string `json:"field"`
		Err   error  `json:"error"`
	}
)

func (e *FieldError) Error() string {
	return fmt.Sprintf("validation has failed for field %s: %s", e.Field, e.Err)
}

func (e *ValidationErrors) Error() string {
	return fmt.Sprintf("validation has failed: %v", *e)
}
