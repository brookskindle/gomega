package asyncassertion_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestAsyncAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AsyncAssertion Suite")
}
