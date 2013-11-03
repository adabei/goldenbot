// Package stats logs user stats
package stats

import (
	"fmt"
	"github.com/adabei/goldenbot/events"
	"github.com/adabei/goldenbot/events/cod4"
	"github.com/adabei/goldenbot/rcon/cod"
)

type Stats struct {
	requests chan rcon.RCONRequest
	events   chan interface{}
}

func NewStats(requests chan rcon.RCONRequest, ea events.Aggregator) *Stats {
	s := new(Stats)
	s.requests = requests
	s.events = ea.Subscribe(s)
	return s
}

func (s *Stats) Start() {
	for in := range s.events {
		// Print all received messages to Stdout
		fmt.Println(in)
	}
}
