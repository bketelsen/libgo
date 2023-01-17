package main

import (
	"fmt"
	"time"

	"github.com/bketelsen/libgo/events"
	"github.com/bketelsen/libgo/log"

	"github.com/bketelsen/libgo/log/stderr"
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
func ExamplePrintln() {
	log := stderr.New()

	log.Println("a simple log message without prefix")
}

type level struct{}

func ExampleWithValue() {
	log := stderr.New()

	info := log.WithValue(level{}, "INFO")
	info.Println("everything's fine")

	err := log.WithValue(level{}, "ERROR")
	err.Println("everything's not fine")
}

type prefix struct{}

func ExampleWithPrefix() {
	fn := func(log log.Log) {
		// do some work
		log.WithValue(level{}, "INFO").Println("everythings cool")
	}

	log := stderr.New()
	fn(log.WithValue(prefix{}, "important function"))
}

func main() {
	ExamplePrintln()
	ExampleWithValue()
	ExampleWithPrefix()
	dw := &events.Subscriber{
		Handler: EventHandler,
	}
	// Subscribe to an Event
	events.Subscribe(dw)
	e := NewTopicStart("mytopic")
	events.Publish(e)

}
