package benchmarks

import (
	"time"
)

type ProgressMonitor struct {
	lastTime time.Time
	worked   int
	Steps    []BenchmarkResult
}

func (pm *ProgressMonitor) Start() {
	pm.lastTime = time.Now()
	pm.worked = 0
}

func (pm *ProgressMonitor) Step() {
	if pm.Steps == nil {
		pm.Steps = make([]BenchmarkResult, 1)
	}
	if len(pm.Steps) <= pm.worked {
		pm.Steps = append(pm.Steps, BenchmarkResult{})
	}
	step := &pm.Steps[pm.worked]
	now := time.Now()
	step.add(now.Sub(pm.lastTime))

	pm.worked++
	pm.lastTime = now
}
