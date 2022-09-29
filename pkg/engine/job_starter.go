package engine

import "streamwork/pkg/api"

type JobStarter struct {
	queue_size     int
	job            api.Job             // the job to start
	executorList   []ComponentExecutor // list of executors
	connectionList []*Connection       // connections between component executors
}

func NewJobStarter(job api.Job) *JobStarter {
	return &JobStarter{
		queue_size: 64,
		job:        job,
	}
}

func (j *JobStarter) Start() error {
	// set up executors for all the components.

	// all components are created now. Build the connections to connect the components together.

	// start all the processes.

	// start web server
	return nil
}

func (j *JobStarter) SetupComponentExecutor() {
	// start from sources in the job and traverse components to create executors
	for s := range j.job.GetSources() {
		e := NewSourceExecutor(s)
		// for each source, traverse the operations connected to it
		j.executorList = append(j.executorList, e)
		j.traverseComponent(s, e)	
	}
}

func (j *JobStarter) traverseComponent(component api.Component, excecutor ComponentExecutor) {
	// s := component.GetOutgoingStream()

	// for o, _ := range s.GetAppliedOperators() {
	// 	e := NewOperatorExecutor(o)
	// }
}