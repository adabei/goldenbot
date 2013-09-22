package perfrep

import (
	"fmt"
	"strings"
	"github.com/adabei/goldenbot/rcon"
  "github.com/adabei/goldenbot/events"
  "github.com/adabei/goldenbot/events/cod4"
)

type PerfRep struct {
	Show     bool
	requests chan rcon.RCONRequest
	Stats    map[string]Stats
  events   chan interface{}
}

type Stats struct {
	Kills   int
	Deaths  int
	Assists int
}

func NewPerfRep(show bool, requests chan rcon.RCONRequests, ea events.Aggregator) *MVP {
	p := new(PerfRep)
	p.Show = show
	p.requests = requests
  p.events = ea.Subscribe(p)
	return p
}

func (p *PerfRep) Start() {
	for {
		in := <-p.events
	}
}

// Wilson calculates the <top> best players according to
// www.evanmiller.org/how-not-to-sort-by-average-rating.html
func (p *PerfRep) Wilson(top int) []string {
	return []string{"1", "2"}
}

func (s *Stats) KillDeathRatio() float64 {
	return s.Kills / s.Deaths
}
