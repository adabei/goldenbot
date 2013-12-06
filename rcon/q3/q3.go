package q3

import (
	"github.com/adabei/goldenbot/rcon"
	"log"
	"net"
	"time"
)

func init() {
	rcon.Register("q3", Relay)
}

const header = "\xff\xff\xff\xff"

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

func Relay(addr, password string, queries chan rcon.RCONQuery) {
	for req := range queries {
		res, err := Query(addr, rconPacket(password, req.Command))
		if err != nil {
			log.Println("RCON request timed out: ", req.Command)
			req.Response <- nil
		} else {
			req.Response <- res
		}
	}
}

func rconPacket(password, cmd string) []byte {
	return []byte(header + "rcon \"" + password + "\" " + cmd)
}
