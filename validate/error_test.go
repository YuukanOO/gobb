package validate

import (
	"errors"
	"testing"

	"github.com/yuukanoo/gobb/assert"
)

func TestFieldErrorString(t *testing.T) {
	err := &FieldError{
		Field: "firstName",
		Err:   errors.New("required"),
	}

	assert.Equals(t, "validation has failed for field firstName: required", err.Error(), "error representation should match")
}

func TestValidationErrorsString(t *testing.T) {
	err := &ValidationErrors{
		"firstName": "required",
		"lastName":  "required",
		"avatarUrl": "url",
	}

	assert.Equals(t, "validation has failed: map[avatarUrl:url firstName:required lastName:required]", err.Error(), "error representation should match")
}

func TestErrorsString(t *testing.T) {
	err := Errors{
		"firstName": "required",
		"nickName":  "min:6",
	}

	assert.Equals(t, "validation has failed with 2 errors", err.Error(), "error representation should match")
}
