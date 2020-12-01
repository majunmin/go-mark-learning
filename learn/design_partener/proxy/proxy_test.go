/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package proxy

import "testing"

func TestProxy(t *testing.T) {
	station := &Station{stock: 20}
	sp := NewStationProxy(station)

	station.sell("小名")
	sp.sell("小马")
	station.sell("小花")
	station.sell("小鱼")
}
