package events

import (
	"sync"
	"time"
)

var defaultPublisher = &Publisher{}

// The Event interface represents a type that
// can give information about an event that was
// fired.
type Event interface {
	Name() string
	Created() time.Time
}

// A Subscriber registers handlers for notifications
// of Events.
type Subscriber struct {
	Handler func(Event)
}

// A Publisher contains a list of subscribers who will
// be notified of Events.
type Publisher struct {
	mu          sync.Mutex
	subscribers []*Subscriber
}

// Subscribe adds the given Subscriber to the default Publisher
func Subscribe(s *Subscriber) {
	defaultPublisher.Subscribe(s)
}

// Unsubscribe removes the given Subscriber from the default Publisher
func Unsubscribe(s *Subscriber) {
	defaultPublisher.Unsubscribe(s)
}

// Subscribe adds the given Subscriber to the Publisher's list
// of Subscribers.
func (p *Publisher) Subscribe(s *Subscriber) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscribers = append(p.subscribers, s)
}

// Unsubscribe removes the given Subscriber from the Publisher's list
// of Subscribers.
func (p *Publisher) Unsubscribe(s *Subscriber) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var found bool
	var x int
	var sub *Subscriber
	for x, sub = range p.subscribers {
		if sub == s {
			break
		}

	}
	if found {
		p.subscribers = append(p.subscribers[:x], p.subscribers[x+1:]...)
	}
}

// Publish fires the Handler function for all subscribers,
// passing the event as a parameter.
func (p *Publisher) Publish(e Event) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, sub := range p.subscribers {
		sub.Handler(e)
	}

}

// Publish sends the event to the default publisher.
func Publish(e Event) {
	defaultPublisher.Publish(e)
}
