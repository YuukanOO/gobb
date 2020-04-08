package event

import (
	"reflect"
	"testing"
	"time"
)

func TestLogEmbedded(t *testing.T) {
	u := newUser("john")
	equals(t, 1, u.Version, "should be at version 1 since one event has been applied")
	equals(t, 0, u.OriginVersion(), "should have been created from the version 0")

	evt, ok := u.Dequeue()

	equals(t, true, ok, "should have one event")
	equals(t, "john", evt.Subject, "subject should match")
	equals(t, userCreated{"john"}, evt.Data, "data should match")

	_, ok = u.Dequeue()

	equals(t, false, ok, "no other events should have been stored")
}

func TestLogLifetime(t *testing.T) {
	now := time.Now()
	var log Log

	equals(t, 0, log.Version, "should be at version 0 in the start")
	equals(t, 0, log.OriginVersion(), "should be at origin version 0 in the start")

	log.StoreEvent("john", userCreated{"john"}, now)

	equals(t, 1, log.Version, "should be at version with one event")
	equals(t, 1, len(log.changes), "should have one event")
	equals(t, 0, log.OriginVersion(), "should still be at origin version 0")

	evt, _ := log.Dequeue()

	equals(t, Event{
		Subject:   "john",
		Data:      userCreated{"john"},
		EmittedAt: now,
	}, evt, "event should match entirely")
}

// equals is the same as the one in the assert but since using it will cause
// an import cycle in tests, I have to duplicate it here!
func equals(t *testing.T, expected, actual interface{}, explanation string) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	t.Errorf(`
💬 %s
expected:
	%#v
got:
	%#v`, explanation, expected, actual)
}

// user used to test the Log implementation when embedded. This is how I personnaly
// use the Log to write software with entity updating their state for a reason (ie. an event).
// Usage is fairly simple, you define your entity with everything you need, defines
// an apply method which will update the entity state when given event data and
// inside this method, you just need to make sure to store the event with the
// entity identity.
type user struct {
	name string

	Log
}

func (u *user) apply(events ...Data) {
	for _, e := range events {
		switch m := e.(type) {

		case userCreated:
			u.name = m.name

		default:
			continue
		}

		u.StoreEvent(u.name, e, time.Now())
	}
}

type userCreated struct {
	name string
}

func newUser(name string) *user {
	u := &user{}
	u.apply(userCreated{name})
	return u
}
