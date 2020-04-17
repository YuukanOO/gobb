package validate

import "fmt"

// Errors maps field names to a validation error string.
type Errors map[string]string

func (e *Errors) Error() string {
	return fmt.Sprintf("validation has failed: %v", *e)
}
