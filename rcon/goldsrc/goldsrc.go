package goldsrc

import (
	"github.com/adabei/goldenbot/rcon"
	"github.com/adabei/goldenbot/rcon/q3"
	"log"
)

const header = "\xff\xff\xff\xff"

func init() {
	rcon.Register("goldsrc", Relay)
}

func Relay(addr, password string, queries chan rcon.RCONQuery) {
	for req := range Queries {
		// TODO get challenge from response
		challenge, err := q3.Query(addr, challengePacket())
		if err != nil {
			continue
			// TODO log failure to receive challenge
		}
		challenge = challenge[4:]

		res, err := q3.Query(addr, rconPacket(string(challenge), password, req.Command))
		if err != nil {
			log.Println("RCON request timed out: ", req.Command)
			req.Response <- nil
		} else {
			req.Response <- string(res)
		}
	}
}

func rconPacket(challenge, password, cmd string) []byte {
	return []byte(header + "rcon " + challenge + " \"" + password + "\" " + cmd)
}

func challengePacket() []byte {
	return []byte(header + "challenge rcon\n\x00")
}
