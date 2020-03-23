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
