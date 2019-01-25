package main

import (
	"fmt"
	"time"

	"github.com/bketelsen/libgo/events"
)

type TopicStart struct {
	Topic     string
	StartTime time.Time
}

func (t TopicStart) Name() string {
	return "Started " + t.Topic
}
func (t TopicStart) Created() time.Time {
	return t.StartTime
}

func NewTopicStart(name string) *TopicStart {
	ts := &TopicStart{
		Topic:     name,
		StartTime: time.Now(),
	}
	return ts

}

func EventHandler(e events.Event) {
	switch t := e.(type) {
	case *TopicStart:
		// check the topic and create if needed
		fmt.Println(t.Topic)
	default:
		// we don't care
		fmt.Printf("%v, %T", t, t)
	}

}

func main() {
	dw := &events.Subscriber{
		Handler: EventHandler,
	}
	// Subscribe to an Event
	events.Subscribe(dw)
	e := NewTopicStart("mytopic")
	events.Publish(e)
}
