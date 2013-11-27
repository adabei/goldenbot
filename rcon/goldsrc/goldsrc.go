package goldsrc

import (
  "github.com/adabei/goldenbot/rcon/q3"
)

const header = "\xff\xff\xff\xff"

type RCON struct {
	addr      string
	password  string
	Queries chan RCONQuery
}

func NewRCON(addr, password string, queries chan RCONQuery) *RCON {
	r := new(RCON)
	r.addr = addr
	r.password = password
	r.Queries = queries
	return r
}

type RCONQuery struct {
	Command  string
	Response chan string
}

func (r *RCON) Relay(){
  for req := range r.Queries {
    // TODO get challenge from response
    challenge, err := q3.Query(r.addr, challengePacket())
    if err != nil {
      continue
      // TODO log failure to receive challenge
    }
    challenge = challenge[4:]
    
    res, err := q3.Query(r.addr, rconPacket(string(challenge), r.password, req.Command))
    if err != nil {
      // TODO log timeout
    } else {
      req.Response <- string(res)
    }
  }
}

func rconPacket(challenge, password, cmd string) []byte {
	return []byte(header + "rcon "+ challenge + " \"" + password + "\" " + cmd)
}

func challengePacket() []byte {
  return []byte(header + "challenge rcon\n\x00")
}
