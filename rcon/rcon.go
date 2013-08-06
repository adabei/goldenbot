package rcon

import (
	"fmt"
	"net"
	"os"
	"time"
)

type RCON struct {
	addr      string
	password  string
	incomming chan RCONRequest
}

func NewRCON(addr, password string, incomming chan RCONRequest) *RCON {
	r := new(RCON)
	r.addr = addr
	r.password = password
	r.incomming = incomming
	return r
}

type RCONRequest struct {
	Command  string
	Response chan string
}

func NewRCONRequest(command string, response chan string) *RCONRequest {
	rr := new(RCONRequest)
	rr.Command = command
	rr.Response = response
	return rr
}

func (r *RCON) Relay() {
	for req := range r.incomming {
		udpAddr, err := net.ResolveUDPAddr("udp", r.addr)
		checkError(err)

		conn, err := net.DialUDP("udp", nil, udpAddr)
		checkError(err)

		_, err = conn.Write(rconMessage(r.password, req.Command))
		checkError(err)
		if req.Response != nil {
			var buf [65565]byte
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			n, err := conn.Read(buf[0:])
			checkError(err)
			req.Response <- string(buf[0:n])
		}
	}

}

func rconMessage(password, cmd string) []byte {
	var cod4Header string = "\xff\xff\xff\xff"
	return []byte(cod4Header + "rcon \"" + password + "\" " + cmd)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
