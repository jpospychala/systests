package benchmarks

import (
	"time"
)

type BenchmarkResult struct {
	n     int
	min   time.Duration
	max   time.Duration
	total time.Duration
}

func Benchmark(n int, f func(pm *ProgressMonitor)) BenchmarkResult {
	pm := ProgressMonitor{}
	for i := 0; i < n; i++ {
		pm.start()
		f(&pm)
		pm.finish()
	}

	return pm.latency
}

type ProgressMonitor struct {
	lastTime     time.Time
	latency      BenchmarkResult
	responseTime BenchmarkResult
}

func (pm *ProgressMonitor) start() {
	pm.latency.n++
	pm.responseTime.n++
	pm.lastTime = time.Now()
}

func (pm *ProgressMonitor) work(step int) {
	now := time.Now()
	diff := now.Sub(pm.lastTime)
	pm.latency.total += diff

	if pm.latency.min == 0 || pm.latency.min > diff {
		pm.latency.min = diff
	}

	if pm.latency.max < diff {
		pm.latency.max = diff
	}
	pm.lastTime = now
}

func (pm *ProgressMonitor) finish() {
	pm.work(1)
}
