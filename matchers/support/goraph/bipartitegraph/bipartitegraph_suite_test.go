package bipartitegraph_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestBipartitegraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bipartitegraph Suite")
}
