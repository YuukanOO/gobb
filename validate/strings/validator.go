// Package strings provides string validation methods.
package strings

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/yuukanoo/gobb/validate"
)

// TODO should errors be exposed? ¯\_(ツ)_/¯

var (
	errRequired = errors.New("required")
	errURL      = errors.New("url")
)

type (
	// NextFunc provided to validators.
	NextFunc func(string) error

	// ValidatorFunc used to validate a string field.
	ValidatorFunc func(string, NextFunc) error

	errLength struct {
		name     string
		expected int
	}

	errMatch struct {
		field string
	}
)

func (e *errMatch) Error() string {
	return fmt.Sprintf("match:%s", e.field)
}

func (e *errLength) Error() string {
	return fmt.Sprintf("%s:%d", e.name, e.expected)
}

// Field validates that a field pass given validators.
func Field(name string, value string, validators ...ValidatorFunc) *validate.FieldError {
	cur := -1
	count := len(validators)
	var next NextFunc

	next = func(v string) error {
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

// Required checks that the value is not empty.
func Required(value string, next NextFunc) error {
	if empty(value) {
		return errRequired
	}

	return next(value)
}

// Optional called the next validator only if the value is not considered empty.
// If its empty, it will just returns nil without going further.
func Optional(value string, next NextFunc) error {
	if empty(value) {
		return nil
	}

	return next(value)
}

// URL validators will check that the given value is a valid URL.
func URL(value string, next NextFunc) error {
	val, err := url.Parse(value)

	if err != nil || val.Scheme == "" || val.Host == "" {
		return errURL
	}

	return next(value)
}

// MatchField checks that the value match another field value.
func MatchField(field, fieldValue string) ValidatorFunc {
	return func(value string, next NextFunc) error {
		if value != fieldValue {
			return &errMatch{field: field}
		}

		return next(value)
	}
}

// Min checks that a string has at least n characters.
func Min(n int) ValidatorFunc {
	return func(value string, next NextFunc) error {
		if length(value) < n {
			return &errLength{"min", n}
		}

		return next(value)
	}
}

func empty(value string) bool {
	return length(value) == 0
}

func length(value string) int {
	return utf8.RuneCountInString(strings.TrimSpace(value))
}
