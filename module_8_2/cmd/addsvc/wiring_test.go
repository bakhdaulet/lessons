package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bakhdaulet/lessons/module_8_2/pkg/addendpoint"
	"github.com/bakhdaulet/lessons/module_8_2/pkg/addservice"
	"github.com/bakhdaulet/lessons/module_8_2/pkg/addtransport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/discard"
)

func TestHTTP(t *testing.T) {
	svc := addservice.New(log.NewNopLogger(), discard.NewCounter(), discard.NewCounter())
	eps := addendpoint.New(svc, log.NewNopLogger(), discard.NewHistogram())
	mux := addtransport.NewHTTPHandler(eps, log.NewNopLogger())
	srv := httptest.NewServer(mux)
	defer srv.Close()

	for _, testcase := range []struct {
		method, url, body, want string
	}{
		{"GET", srv.URL + "/concat", `{"a":"1","b":"2"}`, `{"v":"12"}`},
		{"GET", srv.URL + "/sum", `{"a":1,"b":2}`, `{"v":3}`},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		if want, have := testcase.want, strings.TrimSpace(string(body)); want != have {
			t.Errorf("%s %s %s: want %q, have %q", testcase.method, testcase.url, testcase.body, want, have)
		}
	}
}
