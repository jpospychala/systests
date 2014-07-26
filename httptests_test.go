package httptests

import (
	"github.com/bmizerany/assert"
	"github.com/jpospychala/systests/benchmarks"
	"testing"
)

func Test_that_request_to_google_completes(t *testing.T) {
	_, err := GetHttpUrl("http://www.google.com", &benchmarks.ProgressMonitor{})
	assert.Equal(t, err, nil)
}

func Test_that_request_to_nowhere_fails(t *testing.T) {
	_, err := GetHttpUrl("http://www.a-domain-that-doesnt-exist", &benchmarks.ProgressMonitor{})
	assert.NotEqual(t, err, nil)
}

func Test_that_request_to_google_has_timing_details(t *testing.T) {
	pm := benchmarks.ProgressMonitor{}
	GetHttpUrl("http://www.google.com", &pm)
	assert.T(t, len(pm.Steps) == 2)
}
