package Testing_practice

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSplit(t *testing.T) {
	Convey("基础用例", t, func() {
		var (
			s = "a:b:c"
			sep = ":"
			expect = []string{"a", "b", "c"}
		)
		got := Split(s, sep)
		So(got, ShouldResemble, expect)
	})
}
