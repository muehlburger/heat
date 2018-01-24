package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestTempHandler(t *testing.T) {
	tt := []struct {
		name  string
		room  string
		value string
		res   int
		err   string
	}{
		{name: "set temp to 22", room: "livingroom", value: "22", res: 22},
		{name: "missing value", room: "livingroom", value: "", err: "missing value"},
		{name: "missing room", room: "", value: "22", err: "missing room"},
		{name: "wrong room", room: "not available", value: "22", err: "missing room"},
		{name: "wrong room and missing value ", room: "wrong", value: "", err: "missing room"},
		{name: "not a number", room: "livingroom", value: "a", err: "not a number: a"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/temp?r="+tc.room+"&v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			tempHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}
			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status 400 Bad Request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status 200 OK; got %v", res.Status)
			}

			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
			if err != nil {
				t.Fatalf("expected an integer; got %v", b)
			}
			if d != tc.res {
				t.Fatalf("expected %v: got %v", tc.res, d)
			}
		})
	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/temp?r=livingroom&v=22", srv.URL))
	if err != nil {
		t.Fatalf("coult not send GET request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("expected an integer; got %s", b)
	}
	if d != 22 {
		t.Fatalf("expected temp to be 22: got %v", d)
	}
}
