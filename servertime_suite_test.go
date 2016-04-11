package servertime

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestServertime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Servertime Suite")
}
