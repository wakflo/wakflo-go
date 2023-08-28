package sdk

// JobProcessStatus enum definition.
type JobProcessStatus string

const (
	Running   JobProcessStatus = "running"
	Pending   JobProcessStatus = "pending"
	Queued    JobProcessStatus = "queued"
	Completed JobProcessStatus = "completed"
	Cancelled JobProcessStatus = "cancelled"
	Failed    JobProcessStatus = "failed"
	Skipped   JobProcessStatus = "skipped"
)

// Values provides list valid values for Enum.
func (JobProcessStatus) Values() (kinds []string) {
	for _, s := range []JobProcessStatus{
		Running,
		Pending,
		Queued,
		Completed,
		Cancelled,
		Failed,
		Skipped,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

// TaskContext .
type TaskContext struct {
	/// Task that owns this schema instance
	PreviousTaskID *string

	/// Task that owns this schema instance
	CurrentTaskID *string

	/// Task that owns this schema instance
	NextTaskID *string

	Steps map[string]TaskStepState

	Env map[string]string

	Input map[string]interface{}

	Job WorkflowInstance

	Step TaskStep

	Workflow Workflow

	Task Task
}

// TaskResult .
type TaskResult struct {
	Output map[string]interface{}
	Errors []SystemActivityLog
}

// SystemActivityLog .
type SystemActivityLog struct {
	Scope     string
	Message   string
	Level     string
	Timestamp int64
}

// TaskStepState .
type TaskStepState struct {
	TaskId string

	Input map[string]interface{}

	Output map[string]interface{}

	Errors []SystemActivityLog

	Status JobProcessStatus
}

// WorkflowInstance .
type WorkflowInstance struct {
	Scope     string
	Message   string
	Level     string
	Timestamp int64
}

// Task .
type Task struct {
	// ID of the Task
	ID string

	// Name of the Task
	Name string

	// Description of the schema
	Description string

	// Description of the schema
	Version string

	// Data .
	Data TaskStepData

	// Properties of the schema
	Properties map[string]interface{}
}

// Workflow .
type Workflow struct {
	ID       string
	Name     string
	Revision int64
}

// TaskStep .
type TaskStep struct {
	// ID of the TaskStep
	ID string

	// Name of the schema
	Label string

	// Data .
	Data TaskStepData
}

// TaskStepData .
type TaskStepData struct {
	// Name of the Task
	Name string

	// Properties of the schema
	Properties map[string]interface{}
}

type PluginFn = func(ctx TaskContext) TaskResult
