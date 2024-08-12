package bgld

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBGLd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BGLd Suite")
}
