package chronos

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"
)

// Jobs is a slice of jobs
type Jobs []Job

// A Job defines a chronos job
// https://github.com/mesos/chronos/blob/master/docs/docs/api.md#job-configuration
type Job struct {
	Name                   string              `json:"name"`
	Command                string              `json:"command"`
	Shell                  bool                `json:"shell,omitempty"`
	Epsilon                string              `json:"epsilon,omitempty"`
	Executor               string              `json:"executor,omitempty"`
	ExecutorFlags          string              `json:"executorFlags,omitempty"`
	Retries                int                 `json:"retries,omitempty"`
	Owner                  string              `json:"owner,omitempty"`
	OwnerName              string              `json:"ownerName,omitempty"`
	Description            string              `json:"description,omitempty"`
	Async                  bool                `json:"async,omitempty"`
	SuccessCount           int                 `json:"successCount,omitempty"`
	ErrorCount             int                 `json:"errorCount,omitempty"`
	LastSuccess            string              `json:"lastSuccess,omitempty"`
	LastError              string              `json:"lastError,omitempty"`
	CPUs                   float32             `json:"cpus,omitempty"`
	Disk                   float32             `json:"disk,omitempty"`
	Mem                    float32             `json:"mem,omitempty"`
	Disabled               bool                `json:"disabled,omitempty"`
	SoftError              bool                `json:"softError,omitempty"`
	DataProcessingJobType  bool                `json:"dataProcessingJobType,omitempty"`
	ErrorsSinceLastSuccess int                 `json:"errorsSinceLastSuccess,omitempty"`
	URIs                   []string            `json:"uris,omitempty"`
	EnvironmentVariables   []map[string]string `json:"environmentVariables,omitempty"`
	Arguments              []string            `json:"arguments,omitempty"`
	HighPriority           bool                `json:"highPriority,omitempty"`
	RunAsUser              string              `json:"runAsUser,omitempty"`
	Container              *Container          `json:"container,omitempty"`
	Schedule               string              `json:"schedule,omitempty"`
	ScheduleTimeZone       string              `json:"scheduleTimeZone,omitempty"`
	Constraints            []map[string]string `json:"constraints,omitempty"`
	Parents                []string            `json:"parents,omitempty"`
}

// FormatSchedule will return a chronos schedule that can be used by the job
// See https://github.com/mesos/chronos/blob/master/docs/docs/api.md#adding-a-scheduled-job for details
// startTime (time.Time): when you want the job to start. A zero time instant means start immediately.
// interval (string): How often to run the job.
// reps (string): How many times to run the job.
func FormatSchedule(startTime time.Time, interval string, reps string) (string, error) {
	if err := validateInterval(interval); err != nil {
		return "", err
	}

	if err := validateReps(reps); err != nil {
		return "", err
	}

	schedule := fmt.Sprintf("%s/%s/%s", reps, formatTimeString(startTime), interval)

	return schedule, nil
}

// RunOnceNowSchedule will return a schedule that starts immediately, runs once,
// and runs every 2 minutes until successful
func RunOnceNowSchedule() string {
	return "R1//P2M"
}

// Jobs gets all jobs that chronos knows about
func (client *Client) Jobs() (*Jobs, error) {
	jobs := new(Jobs)

	err := client.apiGet(ChronosAPIJobs, jobs)

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

// DeleteJob will delete a chronos job
// name: The name of job you wish to delete
func (client *Client) DeleteJob(name string) error {
	return client.apiDelete(path.Join(ChronosAPIJob, name), nil)
}

// DeleteJobTasks will delete all tasks associated with a job.
// name: The name of the job whose tasks you wish to delete
func (client *Client) DeleteJobTasks(name string) error {
	return client.apiDelete(path.Join(ChronosAPIKillJobTask, name), nil)
}

// StartJob can manually start a job
// name: The name of the job to start
// args: A map of arguments to append to the job's command
func (client *Client) StartJob(name string, args map[string]string) error {
	queryValues := url.Values{}
	for key, value := range args {
		queryValues.Add(key, value)
	}

	uri := path.Join(ChronosAPIJob, name) + "?" + queryValues.Encode()
	return client.apiPut(uri, nil)
}

// AddScheduledJob will add a scheduled job
// job: The job you would like to schedule
func (client *Client) AddScheduledJob(job *Job) error {
	return client.apiPost(ChronosAPIAddScheduledJob, job, nil)
}

// AddDependentJob will add a dependent job
func (client *Client) AddDependentJob(job *Job) error {
	return client.apiPost(ChronosAPIAddDependentJob, job, nil)
}

// RunOnceNowJob will add a scheduled job with a schedule generated by RunOnceNowSchedule
func (client *Client) RunOnceNowJob(job *Job) error {
	job.Schedule = RunOnceNowSchedule()
	return client.AddScheduledJob(job)
}

func validateReps(reps string) error {
	if strings.HasPrefix(reps, "R") {
		return nil
	}

	return errors.New("Repetitions string not formatted correctly")
}

func validateInterval(interval string) error {
	if strings.HasPrefix(interval, "P") {
		return nil
	}

	return errors.New("Interval string not formatted correctly")
}

func formatTimeString(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(time.RFC3339Nano)
}
