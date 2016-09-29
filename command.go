package command_metrics

import (
	"time"
)

type commandInterface interface {
	Run() error
	Output() ([]byte, error)
}

// Command is a wraper type of original Command
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

// Run is wrapper command of original Run()
func (proxy *Command) Run() error {
	if Enable {
		startTime := time.Now()
		defer proxy.measure(startTime)
	}
	return proxy.Original.Run()
}

// Output is wrapper command of original Output()
func (proxy *Command) Output() ([]byte, error) {
	if Enable {
		startTime := time.Now()
		defer proxy.measure(startTime)
	}
	return proxy.Original.Output()
}
