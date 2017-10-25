package chronos

// Constants defining the various chronos endpoints
const (
	ChronosAPIJob             = "v1/scheduler/job"
	ChronosAPIJobs            = "v1/scheduler/jobs"
	ChronosAPIKillJobTask     = "v1/scheduler/task/kill"
	ChronosAPIAddScheduledJob = "v1/scheduler/iso8601"
	ChronosAPIAddDependentJob = "v1/scheduler/dependency"
)
