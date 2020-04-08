// Package assert provides common assertions to make tests easier to write / read
// but is very simple and has only a few methods.
package assert

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/yuukanoo/gobb/event"
)

// Equals checks for deep equality using the reflect.DeepEqual methods.
func Equals(t *testing.T, expected, actual interface{}, explanation string) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	error(t, explanation, expected, actual)
}

// Match checks that a given string value match a regular expression.
func Match(t *testing.T, expr *regexp.Regexp, actual string, explanation string) {
	if expr.MatchString(actual) {
		return
	}

	error(t, explanation, expr.String(), actual)
}

// Nil checks that the actual value is equals to nil.
func Nil(t *testing.T, actual interface{}, explanation string) {
	if actual == nil {
		return
	}

	error(t, explanation, nil, actual)
}

// NotNil checks that the actual value is not equals to nil.
func NotNil(t *testing.T, actual interface{}, explanation string) {
	if actual != nil {
		return
	}

	error(t, explanation, "!=nil", actual)
}

// True checks that the actual value is equals to true.
func True(t *testing.T, actual interface{}, explanation string) {
	Equals(t, true, actual, explanation)
}

// False checks that the actual value is equals to false.
func False(t *testing.T, actual interface{}, explanation string) {
	Equals(t, false, actual, explanation)
}

// Events checks that the given Emitter have correctly raise expected events.
// You can omit the event EmittedAt field since it will not be compared to make
// things easier.
//
// This function makes it easy to check how the domain has been modified during
// a call.
func Events(t *testing.T, expected []event.Event, actual event.Emitter) {
	for _, ee := range expected {
		e, ok := actual.Dequeue()

		if !ok {
			error(t, "expected an event", ee, nil)
		}

		Equals(t, ee.Subject, e.Subject, "event subject should match")
		Equals(t, ee.Data, e.Data, "event data should match")
	}

	e, ok := actual.Dequeue()

	if ok {
		Nil(t, e, "no event were expected")
	}
}

func error(t *testing.T, explanation string, expected, actual interface{}) {
	t.Errorf(`
💬 %s
expected:
	%#v
got:
	%#v`, explanation, expected, actual)
}
