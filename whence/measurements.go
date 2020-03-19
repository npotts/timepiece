package whence

import (
	"time"

	"github.com/beevik/ntp"
)

func viaNtp(dial string) *Point {
	resp, err := ntp.Query(dial)
	if err != nil {
		return nil
	}
	return &Point{
		Type:   NTPClient,
		Delta:  resp.ClockOffset,
		Whence: time.Now().Add(resp.ClockOffset).UTC(),
	}
}
