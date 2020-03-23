## Timepiece

This is a relatively simple app that queries a single NTP server and provides some analog clocks


## Building

This project uses [goreleaser](https://goreleaser.com/) to build and deploy this.  To do this:

```sh

    go get -u github.com/npotts/timepiece/...
    cd ${GOPATH:-"${HOME}/go"}/src/github.com/npotts/timepiece
    goreleaser build
    #look in dist/
```

## Usage

There is a rudimentary help:

```sh

√π ./timepiece --help
usage: timepiece [<flags>]

Client to show current NTP time

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
  -h, --ntp="localhost"  NTP dial string to query for time
  -l, --listen=":80"     Listening dial string

```

I build this to run _on_ the NTP server (raspberrypi with PPS & GPS), so I end up running it like `timepiece`, but other may want to `timepiece --ntp 0.pool.ntp.org -l :8080`.


## Routes

If this client is running, you should be able to query `/whence` route via a `GET` you should get some data like the following:

```
curl http://localhost/whence | jq
....

{
  "localtime": {
    "whence": "2020-03-23T04:46:09.791557Z",
    "delta": 0,
    "reference": "localtime"
  },
  "ntp": {
    "whence": "2020-03-23T04:46:09.783162124Z",
    "delta": -12949876,
    "reference": "ntp"
  }
}
```

Now you have a lazy way to get the time via JSON.  Of course there are delays, so dont expect this to be 100% spot on.  I didnt want to calculate all the HTTP overheads  :shurg:
