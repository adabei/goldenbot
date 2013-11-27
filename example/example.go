package example

import (
	"fmt"
	"github.com/adabei/goldenbot/events"
	"github.com/adabei/goldenbot/events/cod4"
	q3 "github.com/adabei/goldenbot/rcon/q3"
)

type Example struct {
	requests chan q3.RCONRequest
	events   chan interface{}
}

func NewExample(requests chan q3.RCONRequest, ea events.Aggregator) *Example {
	e := new(Example)
	e.requests = requests
	e.events = ea.Subscribe(e)
	return e
}

func (e *Example) Start() {
	for in := range e.events {
		// Print all received messages to Stdout
		fmt.Println(in)
	}
}
