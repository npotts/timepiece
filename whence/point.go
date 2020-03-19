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
