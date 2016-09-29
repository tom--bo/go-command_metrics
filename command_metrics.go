package command_metrics

import (
	"time"
)

// Enable is true when logging is on
var Enable = true

// WrapCommand simply wrap originl exec.command
func WrapCommand(name string, command commandInterface) *Command {
	metrics = newMetrics(name)
	proxy = newCommand(command, metrics)
	return proxy
}
