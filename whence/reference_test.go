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
