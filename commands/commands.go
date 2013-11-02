package commands

import (
	"github.com/adabei/goldenbot/events"
	"github.com/adabei/goldenbot/rcon"
	"github.com/adabei/goldenbut/events/cod4"
)

type Commands struct {
	commands map[string]func()
	events   chan interface{}
	requests chan rcon.RCONRequest
}

func NewCommands(commands map[string]func(), requests chan rcon.RCONRequest, ea events.Aggregator) *Commands {
	c := new(Commands)
	c.commands = commands
	c.events = ea.Subscribe(c)
	c.requests = requests
	return c
}

func (c *Commands) Start() {
	for in := range c.events {
		if ev, ok := in.(cod4.Say); ok {
			if cmd, ok := c.commands[ev.Message]; ok {
				cmd()
			}
		}
	}
}
