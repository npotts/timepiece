package whence

import (
	"encoding/json"
	"testing"
)

func TestReference_String(t *testing.T) {
	tests := map[Reference]string{
		LOCALTIME: "localtime",
		NTPClient: "ntp-client",
	}

	for k, v := range tests {
		if k.String() != v {
			t.Error("Didint get right error string")
		}
	}

	var e interface{}
	fxn := func(r Reference) {
		defer func() { e = recover() }()
		_ = r.String()
	}

	if fxn("unknonw"); e == nil {
		t.Error("Exptected an error about invalid Reference")
	}
}

func TestReference_UnmarshallJSON(t *testing.T) {

	var r Reference

	if err := r.UnmarshalJSON(nil); err == nil {
		t.Error("expected an error")
	}

	b := json.RawMessage("\"not a known reference\"")
	if err := r.UnmarshalJSON(b); err == nil {
		t.Error("expected an error")
	}

	b = json.RawMessage("\"localtime\"")
	if err := r.UnmarshalJSON(b); err != nil || r != LOCALTIME {
		t.Error("expected r to be LOCALTIME: ", err, r)
	}
}
