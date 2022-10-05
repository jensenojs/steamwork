package engine

import "streamwork/pkg/api"

/**
 * The executor for operator components. When the executor is started, a new thread
 * is created to call the apply() function of the operator component repeatedly.
 * 
 * Used to inherited by specific operator
 */
type OperatorExecutor struct {
	ComponentExecutorImpl
	operator api.Operator // specific operator, used to execute apply
}

func newOperatorExecutor(op api.Operator) *OperatorExecutor {
	// needs to set or read fields by func
	oe := &OperatorExecutor{
		operator : op,
	}
	oe.setRunOnce(oe.runOnce)
	return oe
}


func (o *OperatorExecutor) Apply(api.Event, []api.Event) error {
	panic("Need to be implemented by specific operator")
}

/* Run process once.
 * @return true if the thread should continue; false if the thread should exist.
 */
func (o *OperatorExecutor) runOnce() bool {
	// read input
	event := o.takeIncomingEvent()
	if event == nil {
		return false
	}

	// apply operatorion
	o.operator.Apply(event, &o.eventCollector)

	// emit out : should work.?
	for _, e := range o.eventCollector {
		o.sendOutgoingEvent(e)
	}

	// clean up event that executed
	o.eventCollector = nil

	return true
}
