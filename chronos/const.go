package chronos

// Constants defining the various chronos endpoints
const (
	ChronosAPIJob             = "scheduler/job"
	ChronosAPIJobs            = "scheduler/jobs"
	ChronosAPIKillJobTask     = "scheduler/task/kill"
	ChronosAPIAddScheduledJob = "scheduler/iso8601"
	ChronosAPIAddDependentJob = "scheduler/dependency"
)
