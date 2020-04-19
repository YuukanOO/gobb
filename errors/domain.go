// Package errors provides a convenient DomainError type used to represents
// an expected error in the application domain.
package errors

import "fmt"

// DomainError represents an error expected by the domain.
// It makes it easy to catch it and to differentiate this error from infrastructure
// errors we should, for example, trigger an internal server error instead of a
// bad request.
type DomainError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Inner   error  `json:"details"`
}

// New instantiates a new DomainError with the given code and message.
// The message is convenient for other developpers to understand what goes wrong.
func New(code, message string) error {
	return &DomainError{
		Code:    code,
		Message: message,
	}
}

// NewWithErr instantiates a new DomainError with the given inner error.
func NewWithErr(code, message string, inner error) error {
	return &DomainError{
		Code:    code,
		Message: message,
		Inner:   inner,
	}
}

func (e *DomainError) Error() string {
	if e.Inner == nil {
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	}

	return fmt.Sprintf(`%s: %s
	%s`, e.Code, e.Message, e.Inner)
}

// Unwrap the DomainError inner error.
func (e *DomainError) Unwrap() error {
	return e.Inner
}

// Extensions transforms a DomainError in a map representation.
// Useful for gqlgen custom errors for example.
func (e *DomainError) Extensions() map[string]interface{} {
	ext := map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}

	if e.Inner != nil {
		ext["details"] = e.Inner
	}

	return ext
}

// First returns the first non-nil error it founds or nil if no errors has been
// given.
func First(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// Any checks if at least one error is non-nil in the given arguments.
func Any(errs ...error) bool {
	return First(errs...) != nil
}
