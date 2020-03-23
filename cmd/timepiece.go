package main

import (
	"os"

	"github.com/alecthomas/kingpin"

	"github.com/npotts/timepiece/api"
)

var (
	app    = kingpin.New("timepiece", `Client to show current NTP time`)
	ntp    = app.Flag("ntp", "NTP dial string to query for time").Short('h').Default("localhost").String()
	listen = app.Flag("listen", "Listening dial string").Short('l').Default(":80").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	s := &api.Server{
		NtpDial: *ntp,
		Listen:  *listen,
	}
	s.Serve()
}
