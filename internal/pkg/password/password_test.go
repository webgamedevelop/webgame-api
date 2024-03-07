package password

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test for password package", func() {
	const raw = "admin12345"

	var hashed []byte

	Context("Test for password package", func() {
		It("Test for Generate", func() {
			var err error
			hashed, err = Generate([]byte(raw))
			Expect(err).Should(Succeed())
		})

		It("Test for Compare", func() {
			var err error
			err = Compare(hashed, []byte(raw))
			Expect(err).Should(Succeed())
		})
	})
})
