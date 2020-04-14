package validate

import (
	"testing"

	"github.com/yuukanoo/gobb/assert"
	"github.com/yuukanoo/gobb/errors"
)

type user struct {
	FirstName string `validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	NickName  string `json:"-" validate:"required,min=6"`
}

func TestValidation(t *testing.T) {
	tests := []struct {
		name  string
		given user
		err   error
	}{
		{
			name:  "no data",
			given: user{},
			err: errors.NewWithErr(ErrCode, ErrMessage, Errors{
				"FirstName": "required",
				"lastName":  "required",
				"NickName":  "required",
			}),
		},
		{
			name:  "with nickname not respecting validation with param",
			given: user{NickName: "joe"},
			err: errors.NewWithErr(ErrCode, ErrMessage, Errors{
				"FirstName": "required",
				"lastName":  "required",
				"NickName":  "min:6",
			}),
		},
		{
			name:  "with valid data",
			given: user{FirstName: "john", LastName: "doe", NickName: "johndoe"},
			err:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Struct(test.given)

			if test.err != nil {
				assert.Equals(t, test.err, err, "errors should match")
				return
			}

			assert.Nil(t, err, "should not have any error")
		})
	}
}
