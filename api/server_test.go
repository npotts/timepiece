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
	if !strings.Contains(rw.Body.String(), `"ntp":null`) {
		t.Error("Body didnt contain ntp parameter")
	}
	s.server.Close()
}
