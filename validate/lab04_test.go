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

	Context("Step 1", func() {
		It("should have a canary.yaml file", func() {
			failMessage = "canary.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../charts/marketsummary/templates/canary.yaml").To(BeAnExistingFile(), failMessage)
		})
/*
		It("should have a valid canary.yaml", func() {
			canaryExpected := expectYamlToParse("../canary.yaml")
			canaryActual := expectYamlToParse("./solution-data/lab01step03/canary.yaml")
			failMessage = "canary.yaml has incorrect configuration\n"
			Expect(canaryActual).To(ValidateYamlObject(canaryExpected), failMessage)
		})
*/
	})

    Context("Step 4", func() {
		It("should have a CurrencyUtils.java file", func() {
			failMessage = "CurrencyUtils.java Doesn't Exist or is in the wrong location\n"
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

func expectYamlToParse(path string) interface{} {
	var output interface{}
	file, err := ioutil.ReadFile(path)
	failMessage := fmt.Sprintf("File at the path, %s, cannot be found. File may be in wrong location or misnamed.\n", path)
	Expect(err).To(BeNil(), failMessage)
	err = yaml.Unmarshal([]byte(file), &output)
	failMessage = fmt.Sprintf("File at the path, %s, could not be parsed as YAML.\n Error: %s\n", path, err)
	Expect(err).To(BeNil(), failMessage)
	return output
}
