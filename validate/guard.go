// Package validate provides validation errors definition and the Guard method
// which should be use to collect field errors returned by subpackages.
package validate

import "github.com/yuukanoo/gobb/errors"

// Guard collects validation results and wraps them in a *DomainError.
// That *DomainError itself wraps a *ValidationErrors which maps each field in
// error with the validation that has failed.
func Guard(results ...*FieldError) error {
	errs := make(ValidationErrors)

	for _, r := range results {
		if r != nil {
			errs[r.Field] = r.Err.Error()
		}
	}

	if len(errs) > 0 {
		return errors.NewWithErr(ErrCode, ErrMessage, &errs)
	}

	return nil
}
