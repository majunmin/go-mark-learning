/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package simple_factory

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloApi_Say(t *testing.T) {

	SetDefaultFailureMode(FailureContinues)
	defer SetDefaultFailureMode(FailureHalts)

	Convey("Equal assertions should be accessible", t, func() {
		api := NewApi(1)
		So("hi majm", ShouldEqual, api.Say("majm"))

		api2 := NewApi(2)
		So("hello majm", ShouldEqual, api2.Say("majm"))
	})
}
