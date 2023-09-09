package sdk

// //go:generate msgp
//go:generate go run github.com/tinylib/msgp
// //go:generate go run github.com/shamaton/msgpackgen

// JobProcessStatus enum definition.
type JobProcessStatus string

const (
	Running   JobProcessStatus = "Running"
	Pending   JobProcessStatus = "Pending"
	Queued    JobProcessStatus = "Queued"
	Completed JobProcessStatus = "Completed"
	Cancelled JobProcessStatus = "Cancelled"
	Failed    JobProcessStatus = "Failed"
	Skipped   JobProcessStatus = "Skipped"
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
	PreviousTaskID *string `json:"previous_task_id" msg:"previous_task_id"`

	/// Task that owns this schema instance
	CurrentTaskID *string `json:"current_task_id" msg:"current_task_id"`

	/// Task that owns this schema instance
	NextTaskID *string `json:"next_task_id" msg:"next_task_id"`

	Steps map[string]TaskStepState `json:"steps" msg:"steps"`

	Env map[string]string `json:"env" msg:"env"`

	Input map[string]interface{} `json:"input" msg:"input"`

	Job WorkflowInstance `json:"job" msg:"job"`

	Step TaskStep `json:"step" msg:"step"`

	Workflow Workflow `json:"workflow" msg:"workflow"`

	Task Task `json:"task" msg:"task"`
}

// TaskResult .
type TaskResult struct {
	Output map[string]interface{} `json:"output" msg:"output"`
	Errors []SystemActivityLog    `json:"errors" msg:"errors"`
}

// SystemActivityLog .
type SystemActivityLog struct {
	Scope     string `json:"scope" msg:"scope"`
	Message   string `json:"message" msg:"message"`
	Level     string `json:"level" msg:"level"`
	Timestamp int64  `json:"timestamp" msg:"timestamp"`
}

// TaskStepState .
type TaskStepState struct {
	TaskId string `json:"task_id" msg:"task_id"`

	Input map[string]interface{} `json:"input" msg:"input"`

	Output map[string]interface{} `json:"output" msg:"output"`

	Errors []SystemActivityLog `json:"errors" msg:"errors"`

	Status JobProcessStatus `json:"status" msg:"status"`
}

// WorkflowInstance .
type WorkflowInstance struct {
	Id         string `json:"id" msg:"id"`
	WorkflowId string `json:"workflow_id" msg:"workflow_id"`
}

// Task .
type Task struct {
	// ID of the Task
	ID string `json:"id" msg:"id"`

	// Name of the Task
	Name string `json:"name" msg:"name"`

	// Description of the schema
	Description string `json:"description" msg:"description"`

	// Description of the schema
	Version string `json:"version" msg:"version"`

	// Data .
	Data TaskStepData `json:"data" msg:"data"`

	// Properties of the schema
	Properties map[string]interface{} `json:"properties" msg:"properties"`
}

// Workflow .
type Workflow struct {
	ID       string `json:"id" msg:"id"`
	Name     string `json:"name" msg:"name"`
	Revision int64  `json:"revision" msg:"revision"`
}

// TaskStep .
type TaskStep struct {
	// ID of the TaskStep
	ID string `json:"id" msg:"id"`

	// Name of the schema
	Label string `json:"label" msg:"label"`

	// Data .
	Data TaskStepData `json:"data" msg:"data"`
}

// TaskProperties .
type TaskProperties struct {
	// Input of the schema
	Input map[string]interface{} `json:"input" msg:"input"`
	// Output of the schema
	Output map[string]interface{} `json:"output" msg:"output"`
	// Authentication of the schema
	Authentication *map[string]interface{} `json:"authentication" msg:"authentication"`
}

// TaskStepData .
type TaskStepData struct {
	// Name of the Task
	Name string `json:"name" msg:"name"`

	// Properties of the schema
	Properties TaskProperties `json:"properties" msg:"properties"`
}

type PluginFn = func(ctx TaskContext) TaskResult
