package optional_test

import (
	o "optional"

	"github.com/wI2L/jettison"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Optional", func() {
	test := struct {
		Text o.Optional[string] `json:"text,omitempty"`
	}{}

	Describe("Check null values when marshaling JSON format", func() {
		BeforeEach(func() {
			test.Text.IsDefined()
		})
		It("Text must be null", func() {
			Expect(jettison.Marshal(test)).To(BeEquivalentTo(`{"text":null}`))
		})
	})

})
