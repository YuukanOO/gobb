// Package times provides time validation methods.
package times

import (
	"errors"
	"time"

	"github.com/yuukanoo/gobb/validate"
)

var (
	errRequired = errors.New("required")
)

type (
	// NextFunc provided to validators.
	NextFunc func(time.Time) error

	// ValidatorFunc used to validate a time field.
	ValidatorFunc func(time.Time, NextFunc) error
)

// Field validates that a field pass given validators.
func Field(name string, value time.Time, validators ...ValidatorFunc) *validate.FieldError {
	cur := -1
	count := len(validators)
	var next NextFunc

	next = func(v time.Time) error {
		cur++
		if cur >= count {
			return nil
		}
		return validators[cur](v, next)
	}

	if err := next(value); err != nil {
		return &validate.FieldError{
			Field: name,
			Err:   err,
		}
	}

	return nil
}

// Required checks that the given value is correctly set.
func Required(value time.Time, next NextFunc) error {
	if value.IsZero() {
		return errRequired
	}

	return next(value)
}
