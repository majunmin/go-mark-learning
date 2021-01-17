/**
  @Author: majm@ushareit.com
  @date: 2021/1/17
  @note:
**/
package container_learn

import (
	"container/list"
	"container/ring"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()
	e0 := l.PushBack(42)
	e1 := l.PushFront(13)
	e2 := l.PushBack(7)

	l.InsertBefore(3, e0)
	l.InsertAfter(196, e1)
	l.InsertAfter(1729, e2)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%#v \n", e.Value.(int))
	}
}

func TestRing(t *testing.T) {

	// 创建一个长度为3 的环
	r := ring.New(3)
	fmt.Printf("ring: %+v\n", *r)

	for i := 0; i < 3; i++ {
		r.Value = i
		r = r.Next()
	}

	fmt.Printf("init ring: %+v\n", *r)

	// sum
	s := 0
	r.Do(func(i interface{}) {
		fmt.Println(i)
		s += i.(int)
	})
	fmt.Printf("sum ring: %d\n", s)

}
