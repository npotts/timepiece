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
