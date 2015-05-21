package chronos_test

import (
	. "github.com/behance/go-chronos/chronos"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var (
		config_stub Config
	)

	BeforeEach(func() {
		config_stub = Config{
			URL:            "fakeurl",
			Debug:          false,
			RequestTimeout: 5,
		}
	})

	Describe("NewClient", func() {
		It("Returns a new client", func() {
			client, err := NewClient(config_stub)

			Expect(client).To(BeAssignableToTypeOf(new(Client)))
			Expect(err).To(BeNil())
		})
	})
})
