package event

import (
	"time"
)

type (
	// Emitter represents an element which can store events.
	Emitter interface {
		// Dequeue the first event of this emitter and returns it. The caller must
		// checks the bool which represents wether or not we have reach the end.
		Dequeue() (Event, bool)
	}

	// Log should be embedded in any types that needs to raise and apply events.
	// It implements the Emitter interface.
	// See the test file to see how it should be used.
	Log struct {
		// Current version of this log. It is exposed to make it easier to persist
		// it to a datastore but it should never be set manually.
		Version int `json:"version"`
		changes []Event
	}
)

// StoreEvent stores an event in the log with given data. It will increments the
// inner Version of the log and append the changes to the inner list.
func (l *Log) StoreEvent(subject Subject, payload Data, emittedAt time.Time) {
	l.changes = append(l.changes, Event{
		Subject:   subject,
		Data:      payload,
		EmittedAt: emittedAt,
	})
	l.Version++
}

// Dequeue implementation for the Log.
func (l *Log) Dequeue() (Event, bool) {
	if len(l.changes) == 0 {
		return Event{}, false
	}

	head, tail := l.changes[0], l.changes[1:]
	l.changes = tail

	return head, true
}

// OriginVersion retrieves the original version on which the current, non dispatched
// events, have been applied. It makes it easy to deal with optimistic locking by
// looking for an element at this version when updating, if no one could be found,
// that means the element has been modified outside this transaction so we should
// abort to prevent data loss.
func (l *Log) OriginVersion() int {
	return l.Version - len(l.changes)
}
