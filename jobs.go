package chronos

type Jobs struct {
	Jobs []Job
}

type Job struct {
	Name                   string              `json:"name",omitempty`
	Command                string              `json:"command",omitempty`
	Shell                  bool                `json:"shell",omitempty`
	Epsilon                string              `json:"epsilon",omitempty`
	Executor               string              `json:"executor",omitempty`
	ExecutorFlags          string              `json:"executorFlags",omitempty`
	Retries                int                 `json:"retries",omitempty`
	Owner                  string              `json:"owner",omitempty`
	OwnerName              string              `json:"ownerName",omitempty`
	Description            string              `json:"description",omitempty`
	Async                  bool                `json:"async",omitempty`
	SuccessCount           int                 `json:"successCount",omitempty`
	ErrorCount             int                 `json:"errorCount",omitempty`
	LastSuccess            string              `json:"lastSuccess",omitempty`
	LastError              string              `json:"lastError",omitempty`
	CPUs                   float32             `json:"cpus",omitempty`
	Disk                   float32             `json:"disk",omitempty`
	Mem                    float32             `json:"mem",omitempty`
	Disabled               bool                `json:"disabled",omitempty`
	SoftError              bool                `json:"softError",omitempty`
	DataProcessingJobType  bool                `json:"dataProcessingJobType",omitempty`
	ErrorsSinceLastSuccess int                 `json:"errorsSinceLastSuccess",omitempty`
	URIs                   []string            `json:"uris",omitempty`
	EnvironmentVariables   []map[string]string `json:"",omitempty`
	Arguments              []string            `json:"arguments",omitempty`
	HighPriority           bool                `json:"highPriority",omitempty`
	RunAsUser              string              `json:"runAsUser",omitempty`
	Container              *Container          `json:"container",omitempty`
	Schedule               string              `json:"schedule",omitempty`
	ScheduleTimeZone       string              `json:"scheduleTimeZone",omitempty`
	Constraints            []map[string]string `json:"constraints",omitempty`
	Parents                []string            `json:"parents",omitempty`
}

// Get all jobs
func (client *Client) Jobs() (*Jobs, error) {
	jobs := new(Jobs)

	if err := client.apiGet(CHRONOS_API_JOBS, &jobs.Jobs); err != nil {
		return nil, err
	} else {
		return jobs, nil
	}
}
