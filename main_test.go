package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("integer number printer", func() {
	It("should return the same input number to output", func() {
		Expect(Diablo(666)).To(Equal(666))
	})
})

func TestGameOfLife(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GameOfLife Suite")
}
