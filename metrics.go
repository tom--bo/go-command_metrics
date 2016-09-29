package command_metrics

import (
	"fmt"
	metrics "github.com/rcrowley/go-metrics"
	"sync"
	"time"
)

// metrics from "github.com/rcrowley/go-metrics"
type Metrics struct {
	name  string
	timer metrics.Timer
}

func newMetrics(name string) *Metrics {
	return &Metrics{
		name:  name,
		timer: metrics.NewTimer(),
	}
}

var mutex = sync.Mutex{}

func (proxy *Metrics) printMetrics(elapsedTime time.Duration) {
	count := proxy.timer.Count()
	if count > 0 {
		fmt.Printf(
			"time:%v\tcommand:%s\tcount:%d\tsum:%d\telapsed(sec):%6.3f\n",
			time.Now(),
			proxy.name,
			proxy.timer.Count(),
			proxy.timer.Sum(),
			float64(elapsedTime)/float64(time.Second),
		)
	}
}

func (proxy *Metrics) measure(startTime time.Time) {
	elapsedTime := time.Now().Sub(startTime)
	mutex.Lock()
	if proxy.timer == nil {
		proxy.timer = metrics.NewTimer()
	}
	mutex.Unlock()
	proxy.timer.Update(elapsedTime)
	proxy.printMetrics(elapsedTime)
}
