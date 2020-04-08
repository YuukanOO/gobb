package event

type (
	// Dispatcher represents a system capable of dispatching domain events from
	// emitters.
	Dispatcher interface {
		// Use register one or many handlers on this dispatcher.
		Use(...Handler)
		// Dispatch one or more events. You should always prefer the DispatchFrom
		// methods but making it easy to dispatch events could be great too.
		Dispatch(...Event)
		// DispatchFrom dequeue events from given emitters and dispatch them
		// to all listeners.
		DispatchFrom(...Emitter)
		// Close should release every resources such as a connection to a remote
		// broker or that kind of stuff.
		Close()
	}

	// Handler which will be called by a dispatcher when events are dispatched.
	Handler func(Event)

	// inProcessDispatcher is a simple dispatcher calling handlers as soon as
	// an event is dispatched.
	inProcessDispatcher struct {
		handlers []Handler
	}
)

// NewInProcessDispatcher constructs a new simple Dispatcher which will call handlers
// as soon as an event is published.
func NewInProcessDispatcher() Dispatcher {
	return &inProcessDispatcher{
		handlers: make([]Handler, 0),
	}
}

func (d *inProcessDispatcher) Use(handlers ...Handler) {
	d.handlers = append(d.handlers, handlers...)
}

func (d *inProcessDispatcher) Dispatch(events ...Event) {
	for _, evt := range events {
		for _, h := range d.handlers {
			go h(evt)
		}
	}
}

func (d *inProcessDispatcher) DispatchFrom(emitters ...Emitter) {
	for _, emitter := range emitters {
		for {
			evt, ok := emitter.Dequeue()

			if !ok {
				break
			}

			d.Dispatch(evt)
		}
	}
}

func (d *inProcessDispatcher) Close() {
	// Nothing to do for this one :)
}
