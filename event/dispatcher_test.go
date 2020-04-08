package event

import (
	"testing"
	"time"
)

func TestInProcessDispatcherDispatch(t *testing.T) {
	c := make(chan Event)
	dispatcher := NewInProcessDispatcher()
	event := Event{
		Subject: "john",
		Data:    userCreated{"john"},
	}

	h := func(evt Event) {
		c <- evt
	}

	dispatcher.Use(h)
	dispatcher.Dispatch(event)

	select {
	case e := <-c:
		equals(t, event, e, "event should have been raised")
	case <-time.After(5 * time.Second):
		t.Error("looks like nothing has been published to the channel")
	}
}

func TestInProcessDispatcherDispatchFrom(t *testing.T) {
	c := make(chan Event)
	dispatcher := NewInProcessDispatcher()
	u := newUser("john")

	h := func(evt Event) {
		c <- evt
	}

	dispatcher.Use(h)
	dispatcher.DispatchFrom(u)

	select {
	case e := <-c:
		equals(t, "john", e.Subject, "event should have the good subject")
		equals(t, userCreated{"john"}, e.Data, "event should have the good payload")
	case <-time.After(5 * time.Second):
		t.Error("looks like nothing has been published to the channel")
	}
}
