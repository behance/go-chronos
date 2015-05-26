package main

import (
	"fmt"
	"time"

	chronos "github.com/behance/go-chronos/chronos"
)

func main() {
	config := chronos.NewDefaultConfig()
	client, err := chronos.NewClient(config)

	if err != nil {
		return
	}

	// Add a scheduled job
	runSchedule, _ := chronos.FormatSchedule(*new(time.Time), "PT2M", "R1")
	container := chronos.Container{
		Type:  "Docker",
		Image: "libmesos/ubuntu",
	}
	newJob := chronos.Job{
		Name:      "myTestJob",
		Command:   "echo 'Hello World'",
		Container: &container,
		Schedule:  runSchedule,
	}

	client.AddScheduledJob(&newJob)

	// Get all current jobs
	jobs, _ := client.Jobs()
	fmt.Println("Current jobs:")
	for _, job := range *jobs {
		fmt.Println("Job Name: ", job.Name)
	}

	// Delete the job
	client.DeleteJob("myTestJob")

	// Get all current jobs
	jobs, _ = client.Jobs()
	fmt.Println("Current jobs:")
	for _, job := range *jobs {
		fmt.Println("Job Name: ", job.Name)
	}

	// To run a job immediately, and only once
	oneTimeJob := chronos.Job{
		Name:      "myOneTimeJob",
		Command:   "echo 'Hello World'",
		Container: &container,
	}
	client.RunOnceNowJob(&oneTimeJob)
	client.DeleteJob("myOneTimeJob")
}
