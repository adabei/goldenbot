// Package cod4 is an extension to package cod.
// It differs by defining a specific type for COD4's changed Kill event.
package cod4

import (
	"github.com/adabei/goldenbot/events/cod"
	"strings"
)

func Parse(line string) interface{} {
  cleanLine = strings.TrimSpace(line)
  values := strings.Split(cleanLine[strings.Index(cleanLine, " ")+1:], ";")

	switch values[0] {
	case "K":
		ret := Kill{}
		ret.GUIDA = values[1]
		ret.NumA, _ = strconv.Atoi(values[2])
		ret.GUIDB = values[4]
		ret.NumB, _ = strconv.Atoi(values[5])
		ret.NameB = values[7]
		ret.Weapon = values[8]
		ret.DamageDealt, _ = strconv.Atoi(values[9])
		ret.MOD = values[10]
		ret.Target = values[11]
		return ret
	default:
		return cod.Parse(line)
	}
}

// The kill event in CoD4 doesn't state the teams.
type Kill struct {
	GUIDA       string
	NumA        int
	NameA       string
	GUIDB       string
	NumB        int
	NameB       string
	Weapon      string
	DamageDealt int
	MOD         string
	Target      string
}
