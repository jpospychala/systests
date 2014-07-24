package benchmarks

func Benchmark(n int, f func(pm *ProgressMonitor)) []BenchmarkResult {
	pm := ProgressMonitor{steps: make([]BenchmarkResult, 0)}
	for i := 0; i < n; i++ {
		pm.start()
		f(&pm)
		pm.nextStep()
	}

	return pm.steps
}
