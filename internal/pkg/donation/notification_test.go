package donation

import "testing"

func TestCheckNotificationString(t *testing.T) {
	stringExaple := "p2p-incoming&1234567&300.00&643&2011-07-01T09:00:00.000+04:00&41001XXXXXXXX" +
		"&false&01234567890ABCDEF01234567890&YM.label.12345"
	originalHash := "a2ee4a9195f4a90e893cff4f62eeba0b662321f9"
	calculatedHash := checkNotificationString(stringExaple)
	if calculatedHash != originalHash {
		t.Error("bad hash")
	}
}
