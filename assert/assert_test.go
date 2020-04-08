package assert

import (
	"regexp"
	"testing"
	"time"

	"github.com/yuukanoo/gobb/event"
)

func TestEquals(t *testing.T) {
	test := &testing.T{}
	Equals(test, 1, 1, "should be equal")

	if test.Failed() {
		t.Error("should not have failed here")
	}

	Equals(test, 1, 2, "should not be equal")

	if !test.Failed() {
		t.Error("should have failed")
	}
}

func TestTrue(t *testing.T) {
	test := &testing.T{}
	True(test, true, "should be true")

	if test.Failed() {
		t.Error("should not have failed here")
	}

	True(test, false, "should not be true")

	if !test.Failed() {
		t.Error("should have failed")
	}
}

func TestFalse(t *testing.T) {
	test := &testing.T{}
	False(test, false, "should be false")

	if test.Failed() {
		t.Error("should not have failed here")
	}

	False(test, true, "should not be false")

	if !test.Failed() {
		t.Error("should have failed")
	}
}

func TestNil(t *testing.T) {
	test := &testing.T{}
	Nil(test, nil, "should be nil")

	if test.Failed() {
		t.Error("should not have failed here")
	}

	Nil(test, 1, "should not be nil")

	if !test.Failed() {
		t.Error("should have failed")
	}
}

func TestNotNil(t *testing.T) {
	test := &testing.T{}
	NotNil(test, 1, "should not be nil")

	if test.Failed() {
		t.Error("should not have failed here")
	}

	NotNil(test, nil, "should be nil")

	if !test.Failed() {
		t.Error("should have failed")
	}
}

func TestMatch(t *testing.T) {
	r := regexp.MustCompile("^a")
	test := &testing.T{}
	Match(test, r, "a message", "should match")

	if test.Failed() {
		t.Error("should not have failed here")
	}

	Match(test, r, "should not", "should not match")

	if !test.Failed() {
		t.Error("should have failed")
	}
}

func TestEvents(t *testing.T) {
	tests := []struct {
		name           string
		events         []event.Event
		log            func() *event.Log
		expectedToFail bool
	}{
		{
			name:           "empty events in both",
			events:         []event.Event{},
			log:            func() *event.Log { return &event.Log{} },
			expectedToFail: false,
		},
		{
			name: "one event expected",
			events: []event.Event{
				event.Event{Subject: "one", Data: "some"},
			},
			log:            func() *event.Log { return &event.Log{} },
			expectedToFail: true,
		},
		{
			name:   "one event but none expected",
			events: []event.Event{},
			log: func() *event.Log {
				l := &event.Log{}
				l.StoreEvent("one", "some", time.Now())
				return l
			},
			expectedToFail: true,
		},
		{
			name: "many events in both",
			events: []event.Event{
				event.Event{Subject: "one", Data: "some"},
				event.Event{Subject: "two", Data: "other"},
			},
			log: func() *event.Log {
				l := &event.Log{}
				l.StoreEvent("one", "some", time.Now())
				l.StoreEvent("two", "other", time.Now())
				return l
			},
			expectedToFail: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testingT := &testing.T{}

			Events(testingT, test.events, test.log())

			if test.expectedToFail && !testingT.Failed() {
				t.Error("expected to fail but does not")
			}

			if !test.expectedToFail && testingT.Failed() {
				t.Error("does not expect it to fail but it does")
			}
		})
	}
}
