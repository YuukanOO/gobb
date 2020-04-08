// Package durations provides methods to validate durations.
package durations

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
	NextFunc func(time.Duration) error

	// ValidatorFunc used to validate a duration field.
	ValidatorFunc func(time.Duration, NextFunc) error
)

// Field validates that a field pass given validators.
func Field(name string, value time.Duration, validators ...ValidatorFunc) *validate.FieldError {
	cur := -1
	count := len(validators)
	var next NextFunc

	next = func(v time.Duration) error {
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
func Required(value time.Duration, next NextFunc) error {
	if value == 0 {
		return errRequired
	}

	return next(value)
}
