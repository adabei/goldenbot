package "example"

import (
  "fmt"
  "github.com/adabei/goldenbot/rcon"
)

type Example struct {
  requests chan rcon.RCONRequest
}

func NewExample(requests chan rcon.RCONRequest) *Example {
  e := new(Example)
  e.requests = requests
  return e
}

func (e *Example) Start (next, prev chan string) {
  for {
    // Every plugin has to pass on messages to the next
    in := <-prev
    next <- in

    // Here we will print all received messages to Stdout
    fmt.Println(in)
  }
}
