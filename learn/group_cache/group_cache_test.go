/**
  @Author: majm@ushareit.com
  @date: 2021/1/26
  @note:
**/
package group_cache

import "testing"

func TestDemo(t *testing.T) {
	Demo1("127.0.0.1:8002")
}

func TestDemo2(t *testing.T) {
	Demo1("127.0.0.1:8001")
}
