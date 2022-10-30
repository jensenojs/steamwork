package component

import (
	"streamwork/pkg/engine"
	"streamwork/pkg/engine/stream"
)

// ComponentExecutorImpl used to Inherited by OperatorExecutor and SourceExecutor to save the implementation of some methods.
type ComponentExecutorImpl struct {
	Name              string
	Stream            *stream.Stream // connect to next component
	Parallelism       int
	InstanceExecutors []engine.InstanceExecutor
}

type InstanceExecutorImpl struct {
	FnWrapper      func()             // wrapper function for fn
	Fn             func() bool        // process function, need to specific implementation for user logic
	EventCollector []engine.Event     // accept events from user logic
	Incoming       *engine.EventQueue // for upstream processes
	Outgoing       *engine.EventQueue // for downstream processes
}

type EventCollector struct {
	DEFAULT_CHANNEL  engine.Channel
	List             map[engine.Channel][]engine.Event
	RegisterChannels map[engine.Channel]engine.Void
}