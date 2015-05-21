package chronos_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoChronos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoChronos Suite")
}
