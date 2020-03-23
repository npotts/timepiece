// Copyright 2020 Nick Potts

// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
