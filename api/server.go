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

//Package api holds some HTTP API functions
package api

//go:generate packr build

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"

	"github.com/npotts/timepiece/whence"
)

//Server contains enough to start a listening server
type Server struct {
	Listen  string
	NtpDial string

	sampler whence.Sampler
	box     packr.Box
	router  *mux.Router
	server  *http.Server
}

//index serves out the / and variousn reuqest for *.html files
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	b, _ := s.box.Find("index.html")
	w.Write(b)
}

//whence is an actual API call into
func (s *Server) whence(w http.ResponseWriter, r *http.Request) {
	js := json.NewEncoder(w)
	js.Encode(s.sampler.Measure())
}

//Serve will start the instance using the provided parameters
func (s *Server) Serve() error {
	s.sampler = whence.NewSampler(s.NtpDial)
	s.box = packr.NewBox("./templates")
	s.router = mux.NewRouter()
	s.router.HandleFunc("/", s.index)
	s.router.HandleFunc("/*.html", s.index)
	s.router.HandleFunc("/*.htm", s.index)
	s.router.HandleFunc("/whence", s.whence)

	fmt.Println("Starting server on", s.Listen)
	s.server = &http.Server{
		Addr:         s.Listen,
		ReadTimeout:  1 * time.Second,
		IdleTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		Handler:      s.router,
	}
	return s.server.ListenAndServe()
}
