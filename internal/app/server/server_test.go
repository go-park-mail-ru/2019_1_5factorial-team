package server

import (
	"testing"
)

var casesNew = []struct {
	port string
	want string
}{
	{"1", "1"},
	{"", ""},
}

func TestNew(t *testing.T) {
	for _, val := range casesNew {
		s := New(val.port)
		if s.port != val.port {
			t.Error("ERROR expected:", val.port, "have:", s.port)
		}
	}
}
