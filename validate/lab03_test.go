package validate

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 3 Microservices", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 3", func() {
		It("should have a virtualService.yaml file", func() {
			failMessage = "virtualService.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/marketsummary/templates/virtualService.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

    Context("Step 4", func() {
		It("should have a hpa.yaml file", func() {
			failMessage = "hpa.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/marketsummary/templates/hpa.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})
