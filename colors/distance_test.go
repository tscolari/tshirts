package colors_test

import (
	"github.com/tscolari/tshirts/colors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Distance", func() {

	Describe("CalcDistanceHex", func() {
		It("calculates the correct distance between black and white", func() {
			distance := colors.CalcDistanceHex("#000000", "#ffffff")
			Expect(distance).To(Equal(100.0))
		})

		It("calculates the correct distance between same color", func() {
			distance := colors.CalcDistanceHex("#000000", "#ffffff")
			Expect(distance).To(Equal(0.0))
		})
	})
})
