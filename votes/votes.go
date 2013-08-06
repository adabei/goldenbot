package votes

import (
	"fmt"
	"strings"
	"time"
	"github.com/adabei/goldenbot/rcon"
)

type Vote struct {
	Active   bool
	requests chan rcon.RCONRequest
}

func NewVote(requests chan rcon.RCONRequest) *Vote {
	v := new(Vote)
	v.Active = false
	v.requests = requests
	return v
}

func (v *Vote) Start(next, prev chan string) {
	votes := make(map[string]int)
	votef := func(result bool) {}
	end := make(<-chan time.Time)

	for {
		select {
		case in := <-prev:
			next <- in

			if !v.Active {
				if strings.Contains(in, "!votekick") {
					votef = printResults
					v.Active = true
					end = time.After(30000 * time.Millisecond)
				}
			} else {
				if strings.Contains(in, "!yes") {
					votes["fd12ag"] = +1
				} else if strings.Contains(in, "!no") {
					votes["fd12ag"] = -1
				}
			}
		case <-end:
			sum := 0
			for _, value := range votes {
				sum += value
			}
			votes = make(map[string]int)
			v.Active = false
			votef(sum > 0)
		}
	}
}

func printResults(result bool) {
	if result {
		fmt.Println("Vote successful.")
	} else {
		fmt.Println("Vote failed.")
	}
}
