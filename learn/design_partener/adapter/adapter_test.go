/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package adapter

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdapter_Request(t *testing.T) {
	adaptee := &adapteeImpl{}
	adapter := NewAdapter(adaptee)

	Convey("should Equals", t, func() {
		So(adapter.Request(), ShouldEqual, "adaptee method !")
	})
}
