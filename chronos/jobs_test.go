package chronos_test

import (
	"net/http"

	. "github.com/behance/go-chronos/chronos"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	ghttp "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Jobs", func() {
	var (
		config_stub Config
		client      Chronos
		server      *ghttp.Server
	)

	BeforeEach(func() {
		server = ghttp.NewServer()

		config_stub = Config{
			URL:            server.URL(),
			Debug:          false,
			RequestTimeout: 5,
		}

		client, _ = NewClient(config_stub)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Jobs", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/scheduler/jobs"),
					ghttp.RespondWith(http.StatusOK, `[
						{
							"name":"dockerjob",
							"command":"while sleep 10; do date -u +%T; done",
							"shell":true,
							"epsilon":"PT60S",
							"executor":"",
							"executorFlags":"",
							"retries":2,
							"owner":"",
							"ownerName":"",
							"description":"",
							"async":false,
							"successCount":190,
							"errorCount":3,
							"lastSuccess":"2014-03-08T16:57:17.507Z",
							"lastError":"2014-03-01T00:10:15.957Z",
							"cpus":0.5,
							"disk":256.0,
							"mem":512.0,
							"disabled":false,
							"softError":false,
							"dataProcessingJobType":false,
							"errorsSinceLastSuccess":0,
							"uris":[],
							"environmentVariables":[],
							"arguments":[],
							"highPriority":false,
							"runAsUser":"root",
							"container":{
								"type":"docker",
								"image":"libmesos/ubuntu",
								"network":"HOST",
								"volumes":[]
							},
							"schedule":"R/2015-05-21T18:14:00.000Z/PT2M",
							"scheduleTimeZone":""
						}
          ]`),
				),
			)
		})

		It("Makes a request to get all jobs", func() {
			client.Jobs()
			Expect(server.ReceivedRequests()).To(HaveLen(1))
		})

		It("Correctly unmarshalls the response", func() {
			jobs, _ := client.Jobs()
			Expect(jobs).To(Equal(&Jobs{
				Job{
					Name:                 "dockerjob",
					Command:              "while sleep 10; do date -u +%T; done",
					Shell:                true,
					Epsilon:              "PT60S",
					Executor:             "",
					ExecutorFlags:        "",
					Retries:              2,
					Owner:                "",
					Async:                false,
					SuccessCount:         190,
					ErrorCount:           3,
					LastSuccess:          "2014-03-08T16:57:17.507Z",
					LastError:            "2014-03-01T00:10:15.957Z",
					CPUs:                 0.5,
					Disk:                 256,
					Mem:                  512,
					Disabled:             false,
					URIs:                 []string{},
					Schedule:             "R/2015-05-21T18:14:00.000Z/PT2M",
					EnvironmentVariables: []map[string]string{},
					Arguments:            []string{},
					RunAsUser:            "root",
					Container: &Container{
						Type:    "docker",
						Image:   "libmesos/ubuntu",
						Network: "HOST",
						Volumes: []map[string]string{},
					},
				},
			}))
		})
	})

	Describe("DeleteJob", func() {
		var (
			jobName = "fake_job"
		)

		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", "/scheduler/job/"+jobName),
					ghttp.RespondWith(http.StatusOK, nil),
				),
			)
		})

		It("Makes the delete request", func() {
			err := client.DeleteJob(jobName)
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("StartJob", func() {
		var (
			jobName = "fake_job"
		)

		Context("Starting a job with no arguments", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("PUT", "/scheduler/job/"+jobName, ""),
						ghttp.RespondWith(http.StatusOK, nil),
					),
				)
			})

			It("Makes the start request", func() {
				err := client.StartJob(jobName, nil)
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("Starting a job with arguments", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("PUT", "/scheduler/job/"+jobName, "arg1=value1&arg2=value2"),
						ghttp.RespondWith(http.StatusOK, nil),
					),
				)
			})

			It("Can pass arguments to the start job request", func() {
				args := map[string]string{
					"arg1": "value1",
					"arg2": "value2",
				}

				err := client.StartJob(jobName, args)
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
