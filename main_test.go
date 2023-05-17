package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
)

func TestHome(t *testing.T) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	home(rr, r)
	rs := rr.Result() 

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)

	if string(body) != "Hello. This is a site to demo remote docker" {
		t.Error("Wrong home page text")
	}
}