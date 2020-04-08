package validate

import (
	be "errors"
	"testing"

	"github.com/yuukanoo/gobb/assert"
	"github.com/yuukanoo/gobb/errors"
)

func TestGuard(t *testing.T) {
	tests := []struct {
		name     string
		fields   []*FieldError
		expected error
	}{
		{
			name:     "guard with no errors",
			fields:   []*FieldError{nil, nil},
			expected: nil,
		},
		{
			name: "guard with one error",
			fields: []*FieldError{
				&FieldError{Field: "firstName", Err: be.New("required")},
				&FieldError{Field: "avatarUrl", Err: be.New("url")},
			},
			expected: errors.NewWithErr(ErrCode, ErrMessage, &ValidationErrors{
				"firstName": "required",
				"avatarUrl": "url",
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Guard(test.fields...)

			if test.expected == nil {
				assert.Nil(t, err, "no error was expected")
				return
			}

			assert.Equals(t, test.expected, err, "error should match")
		})
	}
}

func TestValidationErrorsAs(t *testing.T) {
	err := Guard(&FieldError{Field: "firstName", Err: be.New("required")},
		&FieldError{Field: "avatarUrl", Err: be.New("url")})

	assert.NotNil(t, err, "should have an error")

	var target *ValidationErrors

	// This is how you typically should checks for ValidationErrors when an error
	// is returned from the domain.

	assert.True(t, be.As(err, &target), "should be able to retrieve a *ValidationErrors")
	assert.Equals(t, &ValidationErrors{
		"firstName": "required",
		"avatarUrl": "url",
	}, target, "should contains field errors")
}
