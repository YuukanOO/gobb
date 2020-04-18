package validate

import (
	"testing"

	"github.com/yuukanoo/gobb/assert"
)

func TestErrorsString(t *testing.T) {
	err := Errors{
		"firstName": "required",
		"nickName":  "min=6",
	}

	assert.Equals(t, "firstName: required, nickName: min=6", err.Error(), "error representation should match")
}
