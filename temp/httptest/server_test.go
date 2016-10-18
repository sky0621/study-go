package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/greet?name=Jiro")
	if err != nil {
		t.Fatalf("http.Get failed: %s", err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read from HTTP Response Body failed: %s", err)
	}

	expected := "Hello, Jiro"
	if string(greeting) != expected {
		t.Fatalf("response of /greet?name=Jiro returns %s, want %s", string(greeting), expected)
	}
}
