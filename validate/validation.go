package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/yuukanoo/gobb/errors"
)

// validate holds the validator instance since it caches stuff.
var validate = validator.New()

const (
	// ErrCode of the validation failed DomainError
	ErrCode = "validation_failed"
	// ErrMessage of the validation failed DomainError
	ErrMessage = "one or more fields are invalid"
)

// Struct validates a struct with struct tags and returns a DomainError if any
// validation has failed. This DomainError wraps an *Errors containing every
// field errors.
//
// Under the hood, it uses go-playground validator v10 with the default struct
// tag.
func Struct(data interface{}) error {
	err := validate.Struct(data)

	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)

	// Conversion failed, that must be something else
	if !ok {
		return errors.NewWithErr(ErrCode, ErrMessage, err)
	}

	// Else translate every errors into something more easy to work with
	errorsMap := make(Errors, len(validationErrors))

	for _, fe := range validationErrors {
		tag := fe.ActualTag()

		if fe.Param() != "" {
			tag += "=" + fe.Param()
		}

		errorsMap[fe.Field()] = tag
	}

	return errors.NewWithErr(ErrCode, ErrMessage, &errorsMap)
}

func init() {
	// Extract the validation field name from the json tag if it exists.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}
