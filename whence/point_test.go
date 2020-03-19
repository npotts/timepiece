package whence

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestPoint_json(t *testing.T) {
	p := &Point{
		Whence: time.Date(2020, 2, 2, 2, 2, 2, 0, time.UTC),
		Delta:  20 * time.Millisecond,
		Type:   LOCALTIME,
	}

	b, e := json.Marshal(p)
	t.Logf("B = %q - %v", string(b), e)

	if e != nil {
		t.Errorf("expected nil error marshalling:, not %v", e)
	}

	rr := &Point{}

	if err := json.Unmarshal(b, rr); err != nil {
		t.Error("Unable to marshal what was decoded into an original struct")
	}

	fmt.Println(string(b))

	if !reflect.DeepEqual(p, rr) {
		t.Logf("B = %q", string(b))
		t.Errorf("Didnt unpacked into the same struct: Got %v / %v", p, rr)
	}
}

func TestSampler_Measure(t *testing.T) {
	s := NewSampler("127.0.0.1")
	m := s.Measure()
	if m == nil {
		t.Error("M shouldnt be nil")
	}
	if _, ok := m[string(LOCALTIME)]; !ok {
		t.Error("There should be a localtime")
	}
	if _, ok := m[string(NTPClient)]; !ok {
		t.Error("There should be a localtime")
	}
}
