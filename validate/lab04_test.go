package validate

import (
	"log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 4 Deployments", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 1", func() {
		It("should have a canary.yaml file", func() {
			failMessage = "canary.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/marketsummary/templates/canary.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

    Context("Step 4", func() {
		It("should have a CurrencyUtils.java file", func() {
			failMessage = "CurrencyUtils.java doesn't exist or is in the wrong location\n"
			Expect("../src/main/java/org/springframework/nanotrader/data/util/CurrencyUtils.java").To(BeAnExistingFile(), failMessage)
		})
	})

	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})
