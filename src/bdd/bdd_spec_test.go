package bdd

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	convey.Convey("give 2 even numbers", t, func() {
		a := 2
		b := 4
		convey.Convey("when add the two numbers", func() {
			c := a + b

			convey.Convey("then the result is still even", func() {
				convey.So(c%2, convey.ShouldEqual, 0)
			})

		})
	})
}
