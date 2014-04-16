// Package source implements the Source RCON protocol.
// This implementation can be used for games using the Source engine,
// most notably Valve games.
package source

//TODO multi-packet responses

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/schwarz/goldenbot/rcon"
	"log"
	"net"
	"time"
)

const (
	SERVERDATA_AUTH           = 3
	SERVERDATA_AUTH_RESPONSE  = 2
	SERVERDATA_EXECCOMMAND    = 2
	SERVERDATA_RESPONSE_VALUE = 0
)

func init() {
	rcon.Register("source", Relay)
}

// Query sends a single RCON command cmd using the connection conn.
// It returns the server response and any error encountered.
func Query(conn net.Conn, packet []byte) ([]byte, error) {
	_, err := conn.Write(packet)
	if err != nil {
		return nil, err
	}

	var buf [4096]byte
	conn.SetReadDeadline(time.Now().Add(5000 * time.Millisecond))
	n, err := conn.Read(buf[0:])
	if err != nil {
		return nil, err
	}

	return buf[12:n], nil
}

// Relay allows for easier use in concurrent system.
// By calling Relay in a goroutine and passing requests using
// the queries channel the password does not need to be exposed.
func Relay(addr, password string, queries chan rcon.RCONQuery) {
	conn, err := authorizeConnection(addr, password)
	if err != nil {
		log.Fatal(err)
	}

	for req := range queries {
		res, err := Query(conn, rconPacket(1337, SERVERDATA_EXECCOMMAND, req.Command))
		if err != nil {
			log.Println("RCON request failed:", req.Command)
		}

		if req.Response != nil {
			req.Response <- res
		}

		time.Sleep(500 * time.Millisecond)
	}
}

// authorizeConnection establishes the RCON connection to the server addr
// using password for authorization.
func authorizeConnection(addr, password string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	id := 0
	conn.Write(rconPacket(id, 3, password))

	var buf [4096]byte

	// serverdata_response_value
	conn.SetReadDeadline(time.Now().Add(5000 * time.Millisecond))
	_, err = conn.Read(buf[0:])
	if err != nil {
		return nil, errors.New("Authorization failed: " + err.Error())
	}

	// serverdata_auth_response
	_, err = conn.Read(buf[0:])
	conn.SetReadDeadline(time.Now().Add(15000 * time.Millisecond))
	if err != nil {
		return nil, errors.New("Authorization failed: " + err.Error())
	}

	resID, _ := binary.Varint(buf[4:7])
	if int(resID) != id {
		return nil, errors.New("Authorization failed: invalid RCON password")
	}

	return conn, nil
}

// rconPacket generates a Source compatible packet.
func rconPacket(id, packetType int, body string) []byte {
	buf := make([]byte, 0)
	buf = append(buf, littleEndianInt32(len(body)+10)...)
	buf = append(buf, littleEndianInt32(id)...)
	buf = append(buf, littleEndianInt32(packetType)...)
	buf = append(buf, []byte(body+"\x00")...)
	buf = append(buf, []byte("\x00")...)

	return buf
}

// littleEndianInt32 converts a 32 bit integer value to []byte.
func littleEndianInt32(n int) []byte {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, int32(n))
	return buf.Bytes()
}
