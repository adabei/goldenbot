// Package cod can be used to parse server log lines generated
// by COD, COD2 and COD4. Note that COD4 differs slightly.
package cod

import (
	"log"
	"strconv"
	"strings"
)

// Parse parses a log line and returns a matching event struct.
// In the case that a line doesn't match anything nil is returned.
func Parse(line string) interface{} {
	indices := map[string]int{"say": 4, "sayteam": 4, "tell": 7}

	offset := 0
	if line[:1] == " " {
		offset = strings.Index(line[1:], " ") + 2
	} else {
		offset = strings.Index(line, " ") + 1
	}

	values := strings.Split(line[offset:], ";")
	switch values[0] {
	case "J":
		ret := Join{}
		ret.GUID = values[1]
		ret.Num, _ = strconv.Atoi(values[2])
		ret.Name = values[3]
		return ret
	case "Q":
		ret := Quit{}
		ret.GUID = values[1]
		ret.Num, _ = strconv.Atoi(values[2])
		ret.Name = values[3]
		return ret
	case "say":
		ret := Say{}
		ret.GUID = values[1]
		ret.Num, _ = strconv.Atoi(values[2])
		ret.Name = values[3]
		ret.Message = strings.Join(values[indices["say"]:], "")
		return ret
	case "sayteam":
		ret := SayTeam{}
		ret.GUID = values[1]
		ret.Num, _ = strconv.Atoi(values[2])
		ret.Name = values[3]
		ret.Message = strings.Join(values[indices["sayteam"]:], "")
		return ret
	case "tell":
		ret := Tell{}
		ret.GUIDA = values[1]
		ret.NumA, _ = strconv.Atoi(values[2])
		ret.NameA = values[3]
		ret.GUIDB = values[4]
		ret.NumB, _ = strconv.Atoi(values[5])
		ret.NameB = values[6]
		ret.Message = strings.Join(values[indices["tell"]:], "")
		return ret
	case "D":
		if values[6] != "-1" {
			ret := Damage{}
			ret.GUIDA = values[1]
			ret.NumA, _ = strconv.Atoi(values[2])
			ret.TeamA = values[3]
			ret.NameA = values[4]
			ret.GUIDB = values[5]
			ret.NumB, _ = strconv.Atoi(values[6])
			ret.TeamB = values[7]
			ret.NameB = values[8]
			ret.Weapon = values[9]
			ret.DamageDealt, _ = strconv.Atoi(values[10])
			ret.MOD = values[11]
			ret.Target = values[12]
			return ret
		} else {
			ret := WorldDamage{}
			ret.GUID = values[1]
			ret.Num, _ = strconv.Atoi(values[2])
			ret.Team = values[3]
			ret.Name = values[4]
			ret.DamageDealt, _ = strconv.Atoi(values[10])
			ret.MOD = values[11]
			return ret
		}
	case "K":
		ret := Kill{}
		ret.GUIDA = values[1]
		ret.NumA, _ = strconv.Atoi(values[2])
		ret.TeamA = values[3]
		ret.NameA = values[4]
		ret.GUIDB = values[5]
		ret.NumB, _ = strconv.Atoi(values[6])
		ret.TeamB = values[7]
		ret.NameB = values[8]
		ret.Weapon = values[9]
		ret.DamageDealt, _ = strconv.Atoi(values[10])
		ret.MOD = values[11]
		ret.Target = values[12]
		return ret
	case "Weapon":
		ret := Weapon{}
		ret.GUID = values[1]
		ret.Num, _ = strconv.Atoi(values[2])
		ret.Name = values[3]
		ret.Pickup = values[4]
		return ret
	case "A":
		ret := Action{}
		ret.GUID = values[1]
		ret.Num, _ = strconv.Atoi(values[2])
		ret.Name = values[3]
		ret.Action = values[4]
		return ret
	default:
		// TODO Info level
		log.Println("Could not parse line: ", line)
		return nil
	}
}

type InitGame struct {
	Unix int64
	Vars map[string]string
}

type ExitLevel struct {
	Unix int64
}

type ShutdownGame struct {
	Unix int64
}

type Weapon struct {
	GUID   string
	Num    int
	Name   string
	Pickup string
}

type Action struct {
	GUID   string
	Num    int
	Name   string
	Action string
}

type Say struct {
	GUID    string
	Num     int
	Name    string
	Message string
}

type SayTeam struct {
	GUID    string
	Num     int
	Name    string
	Message string
}

type Tell struct {
	GUIDA   string
	NumA    int
	NameA   string
	GUIDB   string
	NumB    int
	NameB   string
	Message string
}

type Kill struct {
	GUIDA       string
	NumA        int
	TeamA       string
	NameA       string
	GUIDB       string
	NumB        int
	TeamB       string
	NameB       string
	Weapon      string
	DamageDealt int
	MOD         string
	Target      string
}

type Damage struct {
	GUIDA       string
	NumA        int
	TeamA       string
	NameA       string
	GUIDB       string
	NumB        int
	TeamB       string
	NameB       string
	Weapon      string
	DamageDealt int
	MOD         string
	Target      string
}

type WorldDamage struct {
	GUID        string
	Num         int
	Team        string
	Name        string
	DamageDealt int
	MOD         string
}

type Quit struct {
	GUID string
	Num  int
	Name string
}

type Join struct {
	GUID string
	Num  int
	Name string
}
