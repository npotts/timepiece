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
	"time"
)

//Point is a point in time, with optional information from a refernce
type Point struct {
	Whence time.Time     `json:"whence"`    //datetime  when this point exists
	Delta  time.Duration `json:"delta"`     //delta from some reference to this time.
	Type   Reference     `json:"reference"` //internal reference for type
}

//Sampler can return some samples from some time services
type Sampler interface {
	Measure() map[string]*Point //Measure performs the sample
}

type sampler struct {
	dial string
}

var _ = Sampler(&sampler{})

//Measure performs the sample
func (s *sampler) Measure() map[string]*Point {
	return map[string]*Point{
		string(LOCALTIME): &Point{
			Whence: time.Now().UTC(),
			Delta:  0,
			Type:   LOCALTIME,
		},
		string(NTPClient): viaNtp(s.dial),
	}
}

//NewSampler returns a sampler that tries to use NTP
func NewSampler(dial string) Sampler {
	return &sampler{dial: dial}
}
