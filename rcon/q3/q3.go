package q3

import (
	"net"
	"time"
)

const header = "\xff\xff\xff\xff"

type RCON struct {
	addr     string
	password string
	Queries  chan RCONQuery
}

type RCONQuery struct {
	Command  string
	Response chan []byte
}

func Query(addr, cmd string) ([]byte, error) {
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
	conn.SetReadDeadline(time.now().Add(5000 * time.Millisecond))
	n, err := conn.Read(buf[0:])
	if err != nil {
		return nil, err
	}

	return buf[0:n], nil
}

func (r *RCON) Relay() {
	for req := range r.Queries {
		res, err := Query(r.addr, rconPacket(r.password, req.Command))
		if err != nil {
			//log timeout
		} else {
			req.Response <- res
		}
	}
}

func rconPacket(password, cmd string) []byte {
	return []byte(header + "rcon \"" + password + "\" " + cmd)
}
