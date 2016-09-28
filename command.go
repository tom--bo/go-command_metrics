package command_metrics

import (
	"time"
)

type commandInterface interface {
	Run() error
	Output() ([]byte, error)
}

type Command struct {
	Original commandInterface
	metrics  *Metrics
}

func newCommand(command commandInterface, metrics *Metrics) *Command {
	return &Command{
		command,
		metrics,
	}
}

func (proxy *Command) measure(startTime time.Time) {
	proxy.metrics.measure(startTime)
}

func (proxy *Command) Run() error {
	if Enable {
		startTime := time.Now()
		defer proxy.measure(startTime)
	}
	return proxy.Original.Run()
}

func (proxy *Command) Output() ([]byte, error) {
	if Enable {
		startTime := time.Now()
		defer proxy.measure(startTime)
	}
	return proxy.Original.Output()
}
