package errors

import (
	"errors"
	"testing"

	"github.com/yuukanoo/gobb/assert"
)

func TestNew(t *testing.T) {
	err := New("a_code", "an explanation")

	domainErr, ok := err.(*DomainError)

	assert.True(t, ok, "should be a pointer to a DomainError")
	assert.Equals(t, "a_code", domainErr.Code, "code should match")
	assert.Equals(t, "an explanation", domainErr.Message, "code should match")
	assert.Nil(t, domainErr.Inner, "it should not have an inner error")
}

func TestNewWithErr(t *testing.T) {
	wrappedErr := errors.New("an inner error")
	err := NewWithErr("a_code", "an explanation", wrappedErr)

	domainErr, ok := err.(*DomainError)

	assert.True(t, ok, "should be a pointer to a DomainError")
	assert.Equals(t, "a_code", domainErr.Code, "code should match")
	assert.Equals(t, "an explanation", domainErr.Message, "code should match")
	assert.Equals(t, wrappedErr, domainErr.Inner, "it should have an inner error")
}

func TestErrorString(t *testing.T) {
	err := New("a_code", "an explanation")
	assert.Equals(t, "a_code: an explanation", err.Error(), "error string should match")

	err = NewWithErr("a_code", "an explanation", errors.New("an inner err"))
	assert.Equals(t, "a_code: an explanation\n\tan inner err", err.Error(), "error string should match")
}

func TestUnwrap(t *testing.T) {
	wrappedErr := errors.New("an inner error")
	err := NewWithErr("a_code", "an explanation", wrappedErr)

	assert.True(t, errors.Is(err, wrappedErr), "should unwrap to the inner error")
}

func TestAs(t *testing.T) {
	wrappedErr := errors.New("an inner error")
	err := NewWithErr("a_code", "an explanation", wrappedErr)

	targetErr := &DomainError{}

	assert.True(t, errors.As(err, &targetErr), "should be convertible to another *DomainError")
	assert.Equals(t, "a_code", targetErr.Code, "code should match")
	assert.Equals(t, "an explanation", targetErr.Message, "code should match")
	assert.Equals(t, wrappedErr, targetErr.Inner, "it should have an inner error")
}

func TestExtensionsWithoutInner(t *testing.T) {
	err := New("a_code", "an explanation").(*DomainError)

	assert.Equals(t, map[string]interface{}{
		"code":    "a_code",
		"message": "an explanation",
	}, err.Extensions(), "extensions dict should match")
}

func TestExtensionsWithInner(t *testing.T) {
	wrappedErr := errors.New("an inner error")
	err := NewWithErr("a_code", "an explanation", wrappedErr).(*DomainError)

	assert.Equals(t, map[string]interface{}{
		"code":    "a_code",
		"message": "an explanation",
		"details": wrappedErr,
	}, err.Extensions(), "extensions dict should match")
}

func TestAny(t *testing.T) {
	assert.Nil(t, Any(nil, nil, nil), "should returns no error")

	err := errors.New("one error")

	assert.Equals(t, err, Any(nil, err), "should returns the err")
	assert.Equals(t, err, Any(err, nil, nil), "should returns the err")
}
