package base32

import (
	"testing"
)

func TestCrockford32(t *testing.T) {
	for i := 1; i < 100; i += 3 {
		s := Encode(uint32(i))
		d := Decode(s)
		if d == 0 {
			t.Errorf("Roundtrip failed for %d -> %s -> %d", i, s, d)
		}
	}
}

func TestInvalid(t *testing.T) {
	cases := []string{
		"",
		"1",
		"0123456",
		"?12345",
		"0?2345",
		"01?345",
		"0123?5",
		"01234?",
		"??????",
	}
	for _, c := range cases {
		d := Decode(c)
		if d != 0 {
			t.Errorf("Case %q should decode to 0", c)
		}
	}
}
