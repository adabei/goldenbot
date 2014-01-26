package cod4

import (
	"github.com/adabei/goldenbot/events/cod"
	"testing"
)

func TestKill(t *testing.T) {
	// Kill in CoD4 doesn't carry team information
	expected := Kill{"160913", 1, "PlayerOne", "270913", 2, "PlayerTwo", "ak47_mp", 100, "MOD_HEAD_SHOT", "head"}
	if parsed, ok := Parse(" 11:55 160913;1;PlayerOne;270913;2;PlayerTwo;ak47_mp;100;MOD_HEAD_SHOT;head"); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}
