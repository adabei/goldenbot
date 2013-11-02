package cod

import (
	"testing"
)

func TestInvalid(t *testing.T) {
	parsed := Parse(" 11:55 nil;160913;10;PlayerName")
	if parsed != nil {
		t.Errorf("Parsed should be nil.")
	}
}

func TestJoin(t *testing.T) {
	expected := Join{"160913", 10, "PlayerName"}
	if parsed, ok := Parse(" 11:55 J;160913;10;PlayerName").(Join); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestQuit(t *testing.T) {
	expected := Quit{"160913", 10, "PlayerName"}
	if parsed, ok := Parse(" 11:55 Q;160913;10;PlayerName").(Quit); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestDamage(t *testing.T) {
	expected := Damage{"160913", 1, "axis", "PlayerOne", "270913", 2, "allies", "PlayerTwo", "ak37_mp", 29, "MOD_RIFLE_BULLET", "torso_upper"}
	if parsed, ok := Parse(" 11:55 D;160913;1;axis;PlayerOne;270913;2;allies;PlayerTwo;ak37_mp;29;MOD_RIFLE_BULLET;torso_upper").(Damage); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestWorldDamage(t *testing.T) {
	expected := WorldDamage{"160913", 10, "axis", "PlayerName", 13, "MOD_FALLING"}
	if parsed, ok := Parse(" 11:55 D;160913;10;axis;PlayerName;;-1;world;;none;13;MOD_FALLING;none").(WorldDamage); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestKill(t *testing.T) {
	expected := Kill{"160913", 1, "axis", "PlayerOne", "270913", 2, "allies", "PlayerTwo", "ak37_mp", 100, "MOD_HEAD_SHOT", "head"}
	if parsed, ok := Parse(" 11:55 K;160913;1;axis;PlayerOne;270913;2;allies;PlayerTwo;ak37_mp;100;MOD_HEAD_SHOT;head").(Kill); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestSay(t *testing.T) {
	expected := Say{"160913", 10, "PlayerName", "saymessage"}
	if parsed, ok := Parse(" 11:55 say;160913;10;PlayerName;saymessage").(Say); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestSayTeam(t *testing.T) {
	expected := SayTeam{"160913", 10, "PlayerName", "rush A"}
	if parsed, ok := Parse(" 11:55 sayteam;160913;10;PlayerName;rush A").(SayTeam); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestTell(t *testing.T) {
	expected := Tell{"160913", 1, "PlayerOne", "270913", 2, "PlayerTwo", "a/s/l?"}
	if parsed, ok := Parse(" 11:55 tell;160913;1;PlayerOne;270913;2;PlayerTwo;a/s/l?").(Tell); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestWeapon(t *testing.T) {
	expected := Weapon{"160913", 10, "PlayerName", "ak37_mp"}
	if parsed, ok := Parse(" 11:55 Weapon;160913;10;PlayerName;ak37_mp").(Weapon); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}

func TestAction(t *testing.T) {
	expected := Action{"160913", 10, "PlayerName", "htf_scored"}
	if parsed, ok := Parse(" 11:55 A;160913;10;PlayerName;htf_scored").(Action); ok {
		if parsed != expected {
			t.Errorf("Parsed and expected values do not match")
		}
	} else {
		t.Errorf("Mismatched types")
	}
}
