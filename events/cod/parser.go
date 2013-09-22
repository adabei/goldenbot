package cod

import (
	"strconv"
	"strings"
)

func Parse(line string) interface{} {

	//todo global, todo more
	indices := map[string]int{"say": 5, "sayteam": 5, "tell": 8}

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
	case "D":
		if values[6] == "-1" {
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
		ret.GUIDB = values[4]
		ret.NumB, _ = strconv.Atoi(values[5])
		ret.TeamB = values[6]
		ret.NameB = values[7]
		ret.Weapon = values[8]
		ret.DamageDealt, _ = strconv.Atoi(values[9])
		ret.MOD = values[10]
		ret.Target = values[11]
		return ret
	case "Weapon":
		ret := Weapon{}
		ret.GUID = values[1]
		ret.Num, _ = strconv.Atoi(values[2])
		ret.Name = values[3]
		ret.Pickup = values[4]
  case "A":
    ret := Action{}
    ret.GUID = values[1]
    ret.Num, _ = strconv.Atoi(values[2])
    ret.Name = values[3]
    ret.Action = values[4]
	default:
		return nil
	}
	return nil
}

type InitGame struct {
}

type ExitLevel struct {
}

type ShutdownGame struct {
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
