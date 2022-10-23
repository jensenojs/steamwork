package engine

import "streamwork/pkg/api"

type SourceInstanceExecutor struct {
	InstanceExecutorImpl
	instanceId int
	source     api.Source
}

func newSourceExecutorInstance(id int, so api.Source) *SourceInstanceExecutor {
	var soi = new(SourceInstanceExecutor)
	soi.instanceId = id
	soi.source = so
	soi.source.SetupInstance(id)
	soi.setRunOnce(soi.runOnce)
	return soi
}

func (s *SourceInstanceExecutor) runOnce() bool {
	// generate events
	s.source.GetEvents(&s.eventCollector)

	// emit out
	for _, e := range s.eventCollector {
		s.sendOutgoingEvent(e)
	}

	// clean up event that executed
	s.eventCollector = nil

	return true
}
