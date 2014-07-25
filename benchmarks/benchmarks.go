package benchmarks

func Benchmark(n int, f func(pm *ProgressMonitor)) []BenchmarkResult {
	pm := ProgressMonitor{}
	for i := 0; i < n; i++ {
		pm.Start()
		f(&pm)
		pm.Step()
	}

	return pm.Steps
}
