// Package q3 implements the Quake 3 RCON protocol.
// This implementation can also be used for COD, COD2 and COD4.
package q3

import (
	"github.com/schwarz/goldenbot/rcon"
	"log"
	"net"
	"time"
)

func init() {
	rcon.Register("q3", Relay)
}

const header = "\xff\xff\xff\xff"

// Query sends a single RCON command cmd to the server at addr.
// It returns the server response and any error encountered.
func Query(addr string, cmd []byte) ([]byte, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	_, err = conn.Write(cmd)
	if err != nil {
		return nil, err
	}

	var buf [1024]byte
	conn.SetReadDeadline(time.Now().Add(5000 * time.Millisecond))
	n, err := conn.Read(buf[0:])
	if err != nil {
		return nil, err
	}

	return buf[0:n], nil
}

// Relay allows for easier use in concurrent system.
// By calling Relay in a goroutine and passing requests using
// the queries channel the password does not need to be exposed.
func Relay(addr, password string, queries chan rcon.RCONQuery) {
	for req := range queries {
		res, err := Query(addr, rconPacket(password, req.Command))
		if err != nil {
			log.Println("RCON request timed out:", req.Command)
		}

		if req.Response != nil {
			req.Response <- res
		}

		// only two RCON commands per second (server only accepts two)
		time.Sleep(500 * time.Millisecond)
	}
}

// rconPacket generates a Q3 compatible packet.
func rconPacket(password, cmd string) []byte {
	return []byte(header + "rcon \"" + password + "\" " + cmd)
}
