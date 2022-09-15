package Testing_practice

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// gomonkey目前m1执行会报错。

func TestCompute(t *testing.T) {
	convey.Convey("测试gomonkey网络调用", t, func() {
		//patches := gomonkey.ApplyFunc(networkCompute, func(a, b int) (int, error) {
		//	return 2, nil
		//})
		//defer patches.Reset()

		sum, _ := Compute(1, 1)
		fmt.Println(sum)
		//convey.So(sum, convey.ShouldEqual, 2)
	})


}
