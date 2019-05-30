package random

import "testing"

func TestRandBool(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := RandBool()
		if res != false && res != true {
			t.Error("expected", false, "have", res)
		}
	}
}
