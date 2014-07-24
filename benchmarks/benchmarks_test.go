package benchmarks

import (
	"github.com/bmizerany/assert"
	"testing"
	"time"
)

func Test_that_Benchmark_measures_time(t *testing.T) {
	f := func(pm *ProgressMonitor) {
		time.Sleep(100 * time.Millisecond)
	}
	result := Benchmark(2, f)

	assert.Equal(t, 2, result.n)
	assert.T(t, result.total > 200*time.Millisecond)
	assert.T(t, result.total < 210*time.Millisecond)
	assert.T(t, result.min <= result.max)
	assert.T(t, result.min > 0)
	assert.T(t, result.max > 0)
	assert.T(t, result.min < result.total)
	assert.T(t, result.max < result.total)
	assert.Equal(t, result.total, result.min+result.max)
}
