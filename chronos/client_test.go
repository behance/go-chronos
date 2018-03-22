package chronos_test

import (
	"net/http"

	. "github.com/behance/go-chronos/chronos"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	ghttp "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Client", func() {
	var (
		config_stub Config
		server      *ghttp.Server
	)

	BeforeEach(func() {
		server = ghttp.NewServer()

		config_stub = Config{
			URL:            server.URL(),
			Debug:          false,
			RequestTimeout: 5,
			APIPrefix:      "v1",
		}
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("NewClient", func() {
		It("Returns a new client", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/v1/scheduler/jobs"),
				),
			)

			client, err := NewClient(config_stub)

			Expect(client).To(BeAssignableToTypeOf(new(Client)))
			Expect(err).To(BeNil())
		})

		It("Errors if it cannot hit chronos", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/v1/scheduler/jobs"),
					ghttp.RespondWith(http.StatusInternalServerError, nil),
				),
			)

			_, err := NewClient(config_stub)
			Expect(err).To(MatchError("Could not reach chronos cluster: 500 Internal Server Error"))
		})

		It("Gracefully handles no API prefix", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/scheduler/jobs"),
				),
			)

			config_stub = NewDefaultConfig()
			config_stub.URL = server.URL()

			client, err := NewClient(config_stub)

			Expect(client).To(BeAssignableToTypeOf(new(Client)))
			Expect(err).To(BeNil())
		})
	})
})
