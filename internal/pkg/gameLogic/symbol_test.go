package gameLogic

import "testing"

func TestMatchSymbol(t *testing.T) {
	_, err := MatchSymbol("kek")
	if err == nil {
		t.Error("error expected")
	}

	res, err := MatchSymbol("2")
	if err != nil {
		t.Error("error not expected")
	}

	if res != LR {
		t.Error("expected res:", LR, "have:", res)
	}
}
