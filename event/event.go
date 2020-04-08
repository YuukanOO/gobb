// Package event deals with domain events. I really like to make every mutation
// of an entity as an event dispatched once the entity has been persisted.
// This make it really easy to design extensible and reliable systems.
package event

import "time"

type (
	// Subject represents the subject of an event which is the entity on which an
	// event belongs.
	Subject interface{}

	// Data represents event payload and can be anything.
	Data interface{}

	// Event used to represents an event in a system.
	Event struct {
		// Subject retrieve the subject on which the Data has been applied.
		Subject Subject
		// Data retrieves the event payload as raised by the domain.
		Data Data
		// EmittedAt retrieves the time at which the event has been raised.
		EmittedAt time.Time
	}
)
