package httptests

import (
	"testing"
)

func Test_that_request_to_google_completes(t *testing.T) {
	_, _, err := GetHttpUrl("http://www.google.com")

	if err != nil {
		t.Error(err)
	}
}

func Test_that_request_to_nowhere_fails(t *testing.T) {
	_, _, err := GetHttpUrl("http://www.a-domain-that-doesnt-exist")

	if err == nil {
		t.Error("expected error")
	}
}

func Test_that_request_to_google_has_timing_details(t *testing.T) {
	_, timing, _ := GetHttpUrl("http://www.google.com")

	if timing.latency >= timing.responseTime {
		t.Error("latency cannot be longer or equal to response time")
	}
}
