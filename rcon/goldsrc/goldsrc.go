// Package goldsrc implements the GoldSrc RCON protocol.
// This implementation can be used for Counter Strike 1.6.
// It relies on package q3 for DRY-ness.
package goldsrc

import (
	"github.com/schwarz/goldenbot/rcon"
	"github.com/schwarz/goldenbot/rcon/q3"
	"log"
)

const header = "\xff\xff\xff\xff"

func init() {
	rcon.Register("goldsrc", Relay)
}

// Relay allows for easier use in concurrent system.
// By calling Relay in a goroutine and passing requests using
// the queries channel the password does not need to be exposed.
func Relay(addr, password string, queries chan rcon.RCONQuery) {
	for req := range queries {
		// TODO get challenge from response
		challenge, err := q3.Query(addr, challengePacket())
		if err != nil {
			continue
			// TODO log failure to receive challenge
		}
		challenge = challenge[4:]

		res, err := q3.Query(addr, rconPacket(string(challenge), password, req.Command))

		if err != nil {
			log.Println("RCON request timed out:", req.Command)
		}

		if req.Response != nil {
			req.Response <- res
		}
	}
}

// rconPacket generates a GoldSrc compatible packet.
func rconPacket(challenge, password, cmd string) []byte {
	return []byte(header + "rcon " + challenge + " \"" + password + "\" " + cmd)
}

// challengePacket generates a GoldSrc compatible challenge packet.
// Challenge packets are used to receive a challenge nonce.
func challengePacket() []byte {
	return []byte(header + "challenge rcon\n\x00")
}
