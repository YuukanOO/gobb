package durations

import (
	"testing"
	"time"

	"github.com/yuukanoo/gobb/assert"
	"github.com/yuukanoo/gobb/validate"
)

var next = func(time.Duration) error { return nil }

func TestField(t *testing.T) {
	tests := []struct {
		name     string
		given    *validate.FieldError
		expected *validate.FieldError
	}{
		{
			name:  "one validator which fails",
			given: Field("duration", time.Duration(0), Required),
			expected: &validate.FieldError{
				Field: "duration",
				Err:   errRequired,
			},
		},
		{
			name:     "one validator which succeeded",
			given:    Field("duration", 5*time.Hour, Required),
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
	assert.Equals(t, errRequired, Required(time.Duration(0), next), "should be required")
	assert.Nil(t, Required(5*time.Hour, next), "should be good")
}
