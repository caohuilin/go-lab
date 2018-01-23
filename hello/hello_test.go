package hello

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("The value should be greater by one", func() {
		x := 2
		So(x, ShouldEqual, 2)
	})
}
