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

func (br *BenchmarkResult) add(diff time.Duration) {
	br.n++
	br.total += diff

	if br.min == 0 || br.min > diff {
		br.min = diff
	}

	if br.max < diff {
		br.max = diff
	}
}
