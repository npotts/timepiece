package whence

import (
	"testing"
)

func TestViaNtp(t *testing.T) {
	if a := viaNtp("localhost"); a != nil {
		t.Error("Expected a nil response when quering localhost")
	}

	a := viaNtp("0.pool.ntp.org")
	if a == nil {
		t.Error("Got nil  - expected a struct")
		t.FailNow()
	}

	if a.Delta == 0 {
		t.Error("There is no way there is zero delta")
	}

	if a.Type != NTPClient {
		t.Error("Expected NTP type")
	}

	if a.Whence.IsZero() {
		t.Error("Expected non-zero time")
	}
}
