[![Build Status](https://travis-ci.org/behance/go-chronos.svg?branch=master)](https://travis-ci.org/behance/go-chronos)

# go-chronos

Go wrapper for the chronos API.


## Usage
See the examples directory for some more examples.

```go
import (
  chronos "github.com/behance/go-chronos/chronos"
)

// chronos.NewDefaultConfig() provides a quick default config with
// the url set to http://127.0.0.1:4400
config := chronos.Config{
  URL: "http://some-url:4400"
}

client, err := chronos.NewClient(config)

if err != nil {
  //handle however you want
}

jobs, err := client.Jobs()   // To get all jobs chronos knows about
...
```

## Api Calls
These calls all correspond to endpoints described here: https://github.com/mesos/chronos/blob/master/docs/docs/api.md

- Jobs() (*Jobs, error)
- DeleteJob(name string) error
- DeleteJobTasks(name string) error
- StartJob(name string, args map[string]string) error
- AddScheduledJob(job *Job) error
- AddDependentJob(job *Job) error

The following is a simple convenience function:
- RunOnceNowJob(job *Job) error - This will schedule a job to start immediately and run only one time
