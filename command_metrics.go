package command_metrics

import (
	"time"
)

var Enable = true

var proxyRegistry = make(map[commandInterface](*Command))
var metricsRegistry = make(map[string](*Metrics))

func WrapCommand(name string, command commandInterface) *Command {
	metrics := metricsRegistry[name]
	if metrics == nil {
		metrics = newMetrics(name)
		metricsRegistry[name] = metrics
	}
	proxy := proxyRegistry[command]
	if proxy == nil {
		proxy = newCommand(command, metrics)
		proxyRegistry[command] = proxy
	}
	return proxy
}

func Flush() {
	for _, metrics := range metricsRegistry {
		metrics.printMetrics(-1)
	}
}

func Print(duration int) {
	if duration <= 0 {
		return
	}
	timeDuration := time.Duration(duration)
	go func() {
		time.Sleep(timeDuration * time.Second)
		for {
			startTime := time.Now()
			//for _, metrics := range metricsRegistry {
			//	metrics.printMetrics(duration)
			//}
			elapsedTime := time.Now().Sub(startTime)
			time.Sleep(timeDuration*time.Second - elapsedTime*time.Second)
		}
	}()
}
