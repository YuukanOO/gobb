package strings

import (
	"errors"
	"testing"

	"github.com/yuukanoo/gobb/assert"
	"github.com/yuukanoo/gobb/validate"
)

var next = func(string) error { return nil }

func TestField(t *testing.T) {
	tests := []struct {
		name     string
		given    *validate.FieldError
		expected *validate.FieldError
	}{
		{
			name:  "one validator which fails",
			given: Field("firstName", "", Required),
			expected: &validate.FieldError{
				Field: "firstName",
				Err:   errRequired,
			},
		},
		{
			name:     "one validator which succeeded",
			given:    Field("firstName", "john", Required),
			expected: nil,
		},
		{
			name:  "multiple validators which fails",
			given: Field("firstName", "john", Required, URL),
			expected: &validate.FieldError{
				Field: "firstName",
				Err:   errURL,
			},
		},
		{
			name:     "multiple validators which succeed",
			given:    Field("firstName", "https://somewhere.else", Required, URL),
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equals(t, test.expected, test.given, "field error should match")
		})
	}
}

func TestErrMatchString(t *testing.T) {
	assert.Equals(t, "match:password", (&errMatch{"password"}).Error(), "error string should match")
}

func TestErrLengthString(t *testing.T) {
	assert.Equals(t, "min:6", (&errLength{"min", 6}).Error(), "error string should match")
}

func TestRequired(t *testing.T) {
	assert.Equals(t, errRequired, Required("", next), "should be required")
	assert.Equals(t, errRequired, Required("      ", next), "should be required when only whitespaces")
	assert.Nil(t, Required("john", next), "should be good")
}

func TestURL(t *testing.T) {
	assert.Equals(t, errURL, URL("", next), "empty should not be considered as a valid URL")
	assert.Equals(t, errURL, URL("http:", next), "should not be considered as a valid URL")
	assert.Nil(t, URL("http://some.address/else/where", next), "should be good")
}

func TestOptional(t *testing.T) {
	err := errors.New("validator error")
	n := func(value string) error {
		return err
	}

	assert.Equals(t, err, Optional("john", n), "should have call the next validator")
	assert.Nil(t, Optional("", n), "should not call the next validator if the value is empty")
}

func TestMin(t *testing.T) {
	assert.Equals(t, &errLength{"min", 5}, Min(5)("", next), "should returns an err length for empty strings")
	assert.Equals(t, &errLength{"min", 5}, Min(5)("joe", next), "should returns an err length")
	assert.Nil(t, Min(5)("johnn", next), "should be good now")
}

func TestMatchField(t *testing.T) {
	assert.Equals(t, &errMatch{"password"}, MatchField("password", "test")("joe", next), "should returns an err match")
	assert.Nil(t, MatchField("password", "test")("test", next), "should be good now")
}
