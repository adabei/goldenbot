// Package rcon provides general interfaces and structs, as well as helper functions.
package rcon

import (
	"log"
)

type RCONQuery struct {
	Command  string
	Response chan []byte
}

var protocols = make(map[string]func(string, string, chan RCONQuery))

func Register(name string, protocol func(string, string, chan RCONQuery)) {
	if _, dup := protocols[name]; dup {
		log.Fatal("Can't register two different constructors for the same RCON protocol.")
	}
	protocols[name] = protocol
}

func Relay(name, addr, password string, queries chan RCONQuery) {
	if r, ok := protocols[name]; ok {
		r(addr, password, queries)
	}
}

// EasyQuery returns a function to easily query via RCON.
// ch is the chan on which a RCON implementation listens for new queries.
func EasyQuery(ch chan RCONQuery) func(string) []byte {
	return func(cmd string) []byte {
		res := make(chan []byte)
		ch <- RCONQuery{Command: cmd, Response: res}

		return <-res
	}
}
