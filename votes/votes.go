package votes

import (
	"fmt"
	"github.com/adabei/goldenbot/events"
	"github.com/adabei/goldenbot/rcon"
	"github.com/adabei/goldenbut/events/cod4"
	"strings"
	"time"
)

type Vote struct {
	Active   bool
	Votes    map[string]func(bool, string)
	requests chan rcon.RCONRequest
	events   chan interface{}
}

func NewVote(votes map[string]func(bool, string), requests chan rcon.RCONRequest, ea events.Aggregator) *Vote {
	v := new(Vote)
	v.Active = false
	v.Votes = votes
	v.requests = requests
	v.events = ea.Subscribe(v)
	return v
}

func (v *Vote) Start() {
	votes := make(map[string]int)
	votef := func(result bool, arg string) {}
	end := make(<-chan time.Time)

	for {
		select {
		case in := <-v.events:
			if in.(type) == cod4.Say {
				ev := cod4.Say(in)

				if !v.Active {
					if cmd, ok := v.Votes[strings.Split(ev.Message, " ")[0]]; ok {
						votef = cmd
						v.Active = true
						fmt.Println("<Player X> called a vote to <do Y>")
						end = time.After(30000 * time.Millisecond)
					}
				} else {
					if strings.HasPrefix(ev.Message, "!yes") {
						fmt.Println("yes received")
						votes[ev.GUID] = +1
					} else if strings.HasPrefix(ev.Message, "!no") {
						votes[ev.GUID] = -1
					}
				}
			}
		case <-end:
			sum := 0
			for _, value := range votes {
				sum += value
			}
			votes = make(map[string]int)
			v.Active = false
			votef(sum > 0, "namessss")
		}
	}
}

func PrintResults(result bool, arg string) {
	if result {
		fmt.Println("Vote successful.")
	} else {
		fmt.Println("Vote failed.")
	}
}
