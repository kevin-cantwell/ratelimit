package ratelimit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRatelimit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ratelimit Suite")
}
