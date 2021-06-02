package main

// Publisher publishes event to queue.
type Publisher interface {
	Publish(Event) error
}
