package perfrep

import (
  "fmt"
  "strings"
  "github.com/adabei/goldenbot/rcon"
)

type PerfRep struct {
  Show bool
  requests chan rcon.RCONRequest
  Stats map[string]Stats
}

type Stats struct {
  Kills int
  Deaths int
  Assists int
}

func NewPerfRep(show bool, requests chan rcon.RCONRequests) *MVP {
  p := new(PerfRep)
  p.Show = show
  p.requests = requests
  return p
}

func (p *PerfRep) Start(next, prev chan string) {
  for {
    in := <-prev
    next <- in
  }
}

// Wilson calculates the <top> best players according to
// www.evanmiller.org/how-not-to-sort-by-average-rating.html
func (p *PerfRep) Wilson(top int) []string {
  return []string{"1", "2"}
}

func (s *Stats) KillDeathRatio() float64 {
  return s.Kills/s.Deaths
}
