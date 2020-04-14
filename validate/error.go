package validate

import "fmt"

type (
	// Errors maps field names to a validation error string.
	Errors map[string]string

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

func (e Errors) Error() string {
	return fmt.Sprintf("validation has failed with %d errors", len(e))
}
