/**
  @Author: majm@ushareit.com
  @date: 2020/11/28
  @note:
**/
package template

import (
	"testing"
)

func TestLottery_Run(t *testing.T) {
	lottery := &Lottery{}
	lottery.Run(&Context{})
}
