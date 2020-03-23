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

//Package whence contains some tidbits about time measurements and references
package whence

import (
	"encoding/json"
	"fmt"
	"strings"
)

//Reference is a type of clock reference
type Reference string

const (
	//LOCALTIME is a reference from monotonicly increasing system clock
	LOCALTIME Reference = "localtime"

	//NTPClient is a sample reference w.r.t a NTP server
	NTPClient = "ntp"
)

var known = map[Reference]string{
	LOCALTIME: string(LOCALTIME),
	NTPClient: string(NTPClient),
}

//String conforms to Stringer interface
func (r *Reference) String() string {
	if _, ok := known[*r]; ok {
		return string(*r)
	}
	panic(fmt.Errorf("No known type for %q", string(*r)))
}

//MarshalJSON conforms to  json.Marshaler interface
func (r Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

//UnmarshalJSON conforms to  json.Marshaler interface
func (r *Reference) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	s = strings.ToLower(s)
	for k, v := range known {
		if v == s {
			*r = k
			return nil
		}
	}
	return fmt.Errorf("Dont know what sort of reference %q is", s)
}
