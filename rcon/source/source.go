package source

import (
  "github.com/adabei/goldenbot/rcon"
  "log"
  "net"
)

func init() {
  rcon.Register("source", Relay)
}

func Relay(addr, password string, queries chan rcon.RCONRequests) {
  for req := range queries {
  }
}
