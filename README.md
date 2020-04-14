gobb - Go Building Blocks
===

In every projects, there are some patterns and stuff you're always using. This project contains the ones that I always write again and again and as such, I've decided to extract them and refines them here.

Every piece of code here should only depends on the standard library to make it more versatile. You could use one or more packages as you see fit.

Every code in this repo should be tested because it represents the foundation on which I build softwares using the Go language.

*This is my living toolbelt, use it or duplicate it if you want to* 😉

## What does it contains?

- `assert/`: common test assertions to make tests more readable.
- `clock/`: sometimes you need to mock time. This package provide a convenient `clock.Time` which could be easily mocked and returns the current time in UTC. Testing time based stuff is hard and I really like this approach more than any other.
- `errors/`: in every application, you'll have domain specific errors which represents an expected error in the system. This package provides a `DomainError` which you can instantiate using the `errors.New` or `errors.NewWithErr` methods.
- `event/`: provides some stuff to easily write aggregates which could store domain events and dispatchers to forward them to handlers. I really like this approach (see associated tests) because when every state change of an entity is represented as an event, it makes it really easy to extend an application and observe it.
- `logging/`: who doesn't need logging after all? It provides a simple `Logger` interface and a default implementation which writes to given `io.Writer`.
- `validate/`: this one is a WIP, it's hard to find a great way to deal with validation so keep tuned ;)