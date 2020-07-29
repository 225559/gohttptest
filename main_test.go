package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var url string

func TestMain(m *testing.M) {
	srv := httptest.NewServer(handler())
	url = srv.URL
	os.Exit(m.Run())
}

// TestRouting tests whether the patterns return the correct status codes.
func TestRouting(t *testing.T) {
	tt := []struct {
		name       string
		pattern    string
		statusCode int
	}{
		{"empty", "", http.StatusOK},
		{"forwardSlash", "/", http.StatusOK},

		// Do not register this pattern.
		// All unregistered patterns should return StatusNotFound.
		{"abcdefg1234567", "/abcdefg1234567", http.StatusNotFound},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(url + tc.pattern)
			if err != nil {
				t.Fatalf("%v", err)
			}
			if res.StatusCode != tc.statusCode {
				t.Fatalf("expected %v; got %v", tc.statusCode, res.StatusCode)
			}
		})
	}
}
