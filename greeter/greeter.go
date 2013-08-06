// Package greeter provides a barebones mechanic to greet new players.
package greeter

import (
	"fmt"
	"strings"
	"github.com/adabei/goldenbot/rcon"
)

type Greeter struct {
	message  string
	requests chan rcon.RCONRequest
}

func NewGreeter(message string, requests chan rcon.RCONRequest) *Greeter {
	g := new(Greeter)
	g.message = message
	g.requests = requests
	return g
}

func (g *Greeter) Start(next, prev chan string) {
	for {
		in := <-prev
		next <- in

		values := strings.Split(in, " ")
    if strings.HasPrefix(values[1], "J") {
			res := make(chan string)
			g.requests <- *rcon.NewRCONRequest("say \""+fmt.Sprintf(g.message, in)+"\"", res)
		}
	}
}
