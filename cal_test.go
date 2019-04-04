package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerStuff(t *testing.T) {
	Convey("#sum()", t, func() {

		Convey("- Add", func() {
			Convey("Add(sum = (+[]int) +  (+[]int))", nil)
			So(Add(2, 3, 4, 5, 6), ShouldEqual, 20)
			Convey("Add(sum = (-[]int) +  (-[]int))", nil)
			So(Add(-2, -3, -4, -5, -6), ShouldEqual, -20)
		})

		Convey("- Minus", func() {
			Convey("Minus(sum = (+[]int) -  (+[]int))", nil)
			So(Minus(6, 8, 2, 1), ShouldEqual, -5)
			Convey("Minus(sum = (-[]int) -  (-[]int))", nil)
			So(Minus(-2, -3), ShouldEqual, 1)
		})

		Convey("- Multiplies", func() {
			Convey("Multiplies(sum = (+[]int) *  (+[]int))", nil)
			So(Multiplies(2, 3, 4, 5, 6), ShouldEqual, 720)
			Convey("Multiplies(sum = (-[]int) *  (-[]int))", nil)
			So(Multiplies(-2, -3, -4, -5, -6), ShouldEqual, -720)
		})
		Convey("- Divided By", func() {
			Convey("Divided By(sum = []int / []int)", nil)
			a, b := DividedBy(10, 5, 2)
			So(a, ShouldEqual, 1)
			So(b, ShouldEqual, nil)
			c, d := DividedBy(2, 0)
			Convey("Divided By(sum = []int / 0)", nil)
			So(c, ShouldEqual, 0.0)
			So(d, ShouldBeError, errors.New("y != 0"))
		})
	})
}
