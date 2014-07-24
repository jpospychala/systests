package httptests

import (
	"github.com/bmizerany/assert"
	"testing"
)

func Test_that_request_to_google_completes(t *testing.T) {
	_, _, err := GetHttpUrl("http://www.google.com")
	assert.Equal(t, err, nil)
}

func Test_that_request_to_nowhere_fails(t *testing.T) {
	_, _, err := GetHttpUrl("http://www.a-domain-that-doesnt-exist")
	assert.NotEqual(t, err, nil)
}

func Test_that_request_to_google_has_timing_details(t *testing.T) {
	_, timing, _ := GetHttpUrl("http://www.google.com")
	assert.T(t, timing.latency <= timing.responseTime)
}
