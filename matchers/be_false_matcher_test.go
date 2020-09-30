package matchers_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
	. "github.com/brookskindle/gomega/matchers"
)

var _ = Describe("BeFalse", func() {
	It("should handle true and false correctly", func() {
		Expect(true).ShouldNot(BeFalse())
		Expect(false).Should(BeFalse())
	})

	It("should only support booleans", func() {
		success, err := (&BeFalseMatcher{}).Match("foo")
		Expect(success).Should(BeFalse())
		Expect(err).Should(HaveOccurred())
	})
})
