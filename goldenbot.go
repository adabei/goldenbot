package main

import (
  "encoding/json"
	"io/ioutil"
  "os"
	"strings"
	"github.com/adabei/goldenbot/greeter"
	"github.com/adabei/goldenbot/rcon"
	"github.com/adabei/goldenbot/tails"
	"github.com/adabei/goldenbot/votes"
	"github.com/adabei/goldenbot/advert"
)

type GoldenConfig struct {
  Address string
  RCONPassword string
  LogfilePath string
  SayPrefix string
}

func main() {
  fi, _ := os.Open("golden.cfg")
  b, _ := ioutil.ReadAll(fi)
  var cfg GoldenConfig
  json.Unmarshal(b, &cfg)
	rch := make(chan rcon.RCONRequest, 10)
	rcon := rcon.NewRCON(cfg.Address, cfg.RCONPassword, rch)
	greetings := greeter.NewGreeter("Greetings! Welcome to the server, %s.", rch)

	votekick := votes.NewVote(rch)
  advert := advert.NewAdvert("ads.txt", 60000, rch)
	chain := daisy(greetings, votekick, advert)
	go rcon.Relay()

	logch := make(chan string)
	go tails.Tail(cfg.LogfilePath, logch, false)
	for {
		line := <-logch
		//remove go
		func(ch chan string) { ch <- strings.TrimSpace(line) }(chain)
	}
	os.Exit(0)
}

type Plugin interface {
	Start(prev, next chan string)
}

// Daisy sets up the daisy chain of plugins for message passing.
// Returns a channel on which we can send in messages.
func daisy(plugins ...Plugin) chan string {
	last := make(chan string)
	prev := last
	next := last

	for _, p := range plugins {
		prev = make(chan string)
		go p.Start(next, prev)
		next = prev
	}
	// drain the last channel in the chain
	go func(ch chan string) { <-last }(last)
	return prev
}
