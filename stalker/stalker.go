// Package stalker provides a history of all names a user has used and when.
package stalker

import (
	"fmt"
	"github.com/adabei/goldenbot/events"
	"github.com/adabei/goldenbot/events/cod4"
	"github.com/adabei/goldenbot/rcon"
)

type Stalker struct {
	requests chan rcon.RCONRequest
	events   chan interface{}
}

func NewStalker(requests chan rcon.RCONRequest, ea events.Aggregator) *Stalker {
	s := new(Stalker)
	s.requests = requests
	s.events = ea.Subscribe(s)
	return s
}

func (s *Stalker) Start() {
	for in := range s.events {
		// Print all received messages to Stdout
		fmt.Println(in)
	}
}
