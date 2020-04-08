package times

import (
	"testing"
	"time"

	"github.com/yuukanoo/gobb/assert"
	"github.com/yuukanoo/gobb/validate"
)

var next = func(time.Time) error { return nil }

func TestField(t *testing.T) {
	tests := []struct {
		name     string
		given    *validate.FieldError
		expected *validate.FieldError
	}{
		{
			name:  "one validator which fails",
			given: Field("birthDate", time.Time{}, Required),
			expected: &validate.FieldError{
				Field: "birthDate",
				Err:   errRequired,
			},
		},
		{
			name:     "one validator which succeeded",
			given:    Field("birthDate", time.Now(), Required),
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equals(t, test.expected, test.given, "field error should match")
		})
	}
}

func TestRequired(t *testing.T) {
	assert.Equals(t, errRequired, Required(time.Time{}, next), "should be required")
	assert.Nil(t, Required(time.Now(), next), "should be good")
}
