// Package greeter provides a barebones mechanic to greet new players.
package greeter

import (
	"fmt"
	"github.com/adabei/goldenbot/events"
	"github.com/adabei/goldenbot/events/cod"
	rcon "github.com/adabei/goldenbot/rcon/cod"
)

type Greeter struct {
	message  string
	requests chan rcon.RCONRequest
	events   chan interface{}
}

func NewGreeter(message string, requests chan rcon.RCONRequest, ea events.Aggregator) *Greeter {
	g := new(Greeter)
	g.message = message
	g.requests = requests
	g.events = ea.Subscribe(g)
	return g
}

func (g *Greeter) Start() {
	for {
		in := <-g.events
		if ev, ok := in.(cod.Join); ok {
			fmt.Println(ev.Name)
			g.requests <- *rcon.NewRCONRequest("say \""+fmt.Sprintf(g.message, ev.Name)+"\"", nil)
		}
	}
}
