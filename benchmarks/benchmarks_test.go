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
	results := Benchmark(2, f)

	result := results[0]
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

func Test_that_Benchmark_measures_steps_times(t *testing.T) {
	f := func(pm *ProgressMonitor) {
		time.Sleep(100 * time.Millisecond)
		pm.Step()
		time.Sleep(10 * time.Millisecond)
	}
	results := Benchmark(2, f)

	result := results[0]
	assert.Equal(t, 2, result.n)
	assert.T(t, result.total > 200*time.Millisecond)
	assert.T(t, result.total < 210*time.Millisecond)

	result = results[1]
	assert.Equal(t, 2, result.n)
	assert.T(t, result.total > 20*time.Millisecond)
	assert.T(t, result.total < 30*time.Millisecond)
}
