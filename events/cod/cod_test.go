package cod

import (
	"testing"
)

func TestJoin(t *testing.T) {
	if j, ok := Parse("J;160913;10;PlayerName").(Join); ok {
		if j.GUID != "160913" {
			t.Errorf("GUID wrong")
		}

		if j.Num != 10 {
			t.Errorf("Num wrong")
		}

		if j.Name != "PlayerName" {
			t.Errorf("Name wrong")
		}
	} else {
		t.Errorf("Wrong type returned from Parse function")
	}
}

func TestQuit(t *testing.T) {
}
