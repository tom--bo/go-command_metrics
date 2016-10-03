package command_metrics

var Enable = true

var metricsRegistry = make(map[string](*Metrics))

func WrapCommand(name string, command commandInterface) *Command {
	metrics := metricsRegistry[name]
	if metrics == nil {
		metrics = newMetrics(name)
		metricsRegistry[name] = metrics
	}
    proxy := newCommand(command, metrics)
	return proxy
}


