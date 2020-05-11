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
			failMessage = "virtualService.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../charts/marketsummary/templates/virtualService.yaml").To(BeAnExistingFile(), failMessage)
		})
/*
		It("should have a valid virtualService.yaml", func() {
			virtualServiceExpected := expectYamlToParse("../virtualService.yaml")
			virtualServiceActual := expectYamlToParse("./solution-data/lab01step03/virtualService.yaml")
			failMessage = "virtualService.yaml has incorrect configuration\n"
			Expect(virtualServiceActual).To(ValidateYamlObject(virtualServiceExpected), failMessage)
		})
*/
	})

    Context("Step 4", func() {
		It("should have a hpa.yaml file", func() {
			failMessage = "hpa.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../charts/marketsummary/templates/hpa.yaml").To(BeAnExistingFile(), failMessage)
		})
/*
		It("should have a valid hpa.yaml", func() {
			hpaExpected := expectYamlToParse("../hpa.yaml")
			hpaActual := expectYamlToParse("./solution-data/lab01step03/hpa.yaml")
			failMessage = "hpa.yaml has incorrect configuration\n"
			Expect(hpaActual).To(ValidateYamlObject(hpaExpected), failMessage)
		})
*/
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
