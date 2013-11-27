// Package rcon provides general interfaces and structs, as well as helper functions.
package rcon

type RCON interface {
  Relay()
}

type RCONQuery struct {
  Command string
  Response chan []byte
}

// EasyQuery returns a function to easily query via RCON.
// ch is the chan on which a RCON implementation listens for new queries.
func EasyQuery(ch chan RCONQuery) func(string) []byte {
  return func(cmd string) []byte {
    res := make(chan []byte)
    ch <- RCONQuery{Command: cmd, Response: res}

    return <- res
  }
}
