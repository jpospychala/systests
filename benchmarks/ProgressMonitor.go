package benchmarks

import (
	"time"
)

type ProgressMonitor struct {
	lastTime time.Time
	worked   int
	steps    []BenchmarkResult
}

func (pm *ProgressMonitor) start() {
	pm.lastTime = time.Now()
	pm.worked = 0
}

func (pm *ProgressMonitor) nextStep() {
	if len(pm.steps) <= pm.worked {
		pm.steps = append(pm.steps, BenchmarkResult{})
	}
	step := &pm.steps[pm.worked]
	now := time.Now()
	step.add(now.Sub(pm.lastTime))

	pm.worked++
	pm.lastTime = now
}
