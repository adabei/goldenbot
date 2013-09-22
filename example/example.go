package example

import (
  "fmt"
  "github.com/adabei/goldenbot/rcon"
  "github.com/adabei/goldenbot/events"
  "github.com/adabei/goldenbot/events/cod4"
)

type Example struct {
  requests chan rcon.RCONRequest
  events chan interface{}
}

func NewExample(requests chan rcon.RCONRequest, ea events.Aggregator) *Example {
  e := new(Example)
  e.requests = requests
  e.events = ea.Subscribe(e)
  return e
}

func (e *Example) Start () {
  for {
    in := <-e.events

    // Print all received messages to Stdout
    fmt.Println(in)
  }
}
