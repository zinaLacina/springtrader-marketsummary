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

    Context("Step 4", func() {
		It("should have a HasCurrency.java file", func() {
			failMessage = "HasCurrency.java doesn't exist or is in the wrong location\n"
			Expect("../src/main/java/org/springframework/nanotrader/web/controller/HasCurrency.java").To(BeAnExistingFile(), failMessage)
		})
	})

	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})
