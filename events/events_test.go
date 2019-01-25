package events

import (
	"fmt"
	"testing"
	"time"
)

const diskSpaceLowEventName = "DiskSpaceLowEvent"

type DiskSpaceLowEvent struct {
	FreeSpace string
	TimeStamp time.Time
}

func (d DiskSpaceLowEvent) Name() string {
	return diskSpaceLowEventName
}

func (d DiskSpaceLowEvent) Created() time.Time {
	return d.TimeStamp
}

func NewDiskSpaceLowEvent(freespace string) DiskSpaceLowEvent {
	return DiskSpaceLowEvent{
		FreeSpace: freespace,
		TimeStamp: time.Now(),
	}
}

func TestEvents(t *testing.T) {
	var fired bool
	dw := &Subscriber{
		Handler: func(e Event) {
			if dsle, ok := e.(DiskSpaceLowEvent); ok {
				fired = true
				fmt.Println("Disk Space Is Low!", dsle.FreeSpace)
			}
		},
	}
	// Subscribe to an Event
	Subscribe(dw)
	e := DiskSpaceLowEvent{
		FreeSpace: "100G",
		TimeStamp: time.Now(),
	}
	Publish(e)
	if !fired {
		t.Error("Expected event handler to fire")
	}
}
