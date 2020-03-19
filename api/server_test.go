package api

//go:generate packr build

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestServer_Index(t *testing.T) {
	s := &Server{}

	rw := httptest.NewRecorder()
	if s.index(rw, nil); rw.Code != 200 {
		t.Error("Didnt get a good code", rw.Code)
	}
}

func TestServer_whence(t *testing.T) {
	s := &Server{NtpDial: "localhost", Listen: ":8080"}
	go s.Serve()
	time.Sleep(1 * time.Second)

	rw := httptest.NewRecorder()
	if s.whence(rw, nil); rw.Code != 200 {
		t.Error("Expected a 200 error code, not ", rw.Code)
	}

	rw.Body.String()
	if !strings.Contains(rw.Body.String(), "localtime") {
		t.Error("Body didnt contain localtime parameter")
	}
	if !strings.Contains(rw.Body.String(), `"ntp-client":null`) {
		t.Error("Body didnt contain ntp parameter")
	}
	s.server.Close()
}
