package gstruct_test

import (
	"testing"

	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gstruct Suite")
}
