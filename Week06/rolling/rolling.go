package rolling

import (
	"math"
	"sync"
	"time"

	"github.com/Go-000/Week06/window"
)

//RollingCounter ...
type RollingCounter interface {
	Timespan() int
	Add(float64)
	Min() float64
}

type rollingCounter struct {
	mu             sync.RWMutex
	size           int
	w              *window.Window
	lastUpdateTime time.Time
	timeDuration   time.Duration
	index          int
}

//NewRollingCounter ...
func NewRollingCounter(w *window.Window, size int, timeDuration time.Duration) RollingCounter {
	return &rollingCounter{
		w:              w,
		size:           size, //rolling 的时间长度
		lastUpdateTime: time.Now(),
		timeDuration:   timeDuration,
	}
}

func (r *rollingCounter) Timespan() int {
	return int(int(time.Since(r.lastUpdateTime) / r.timeDuration))
}

func (r *rollingCounter) Add(val float64) {
	r.mu.Lock()
	timespan := r.Timespan()
	r.lastUpdateTime = time.Now() //r.lastUpdateTime.Add(time.Duration(timespan * int(r.timeDuration)))
	if timespan > r.size {
		timespan = r.size
	}
	for i := 0; i < r.size; i++ {
		r.w.ResetBucket(i)
		r.index = i
	}
	r.w.Add(r.index+1, val)

}

func (r *rollingCounter) Min() float64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	index := r.index + 1
	if index == r.size {
		index = 0
	}
	result := math.MaxFloat64
	for _, v := range r.w.Widnow[index].Values {
		if result < v {
			result = v
		}
	}
	return result
}
