package commands

import (
  "fmt"
  "strings"
  "github.com/adabei/goldenbot/rcon"
)

type Commands struct {
  commands map[string]func()
  requests chan rcon.RCONRequest
}

func NewCommands(commands map[string]func(), requests chan rcon.RCONRequest) *Commands {
  c := new(Commands)
  c.commands = commands
  c.requests = requests
  return c
}

func (c *Commands) Start (next, prev chan string) {
  for {
    in := <-prev
    next <- in

    if strings.Contains(in, " say;"){
      said := in[strings.LastIndex(in, ";"):]
      if cmd, ok := c.commands[said]; ok {
        cmd()
      }
    }
  }
}
