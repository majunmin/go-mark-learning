/**
  @Author: majm@ushareit.com
  @date: 2020/12/22
  @note:
**/
package sync_learn

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestDemo1(t *testing.T) {
	var optNum int32 = 50

	for i := 0; i < 50; i++ {
		atomic.AddInt32(&optNum, 1)
		time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(time.Second)
	fmt.Println("optNum ", atomic.LoadInt32(&optNum))
}

func TestDemo2(t *testing.T) {

	var num2 int32 = 0
	// 构成简易自旋锁
	// for语句中的 CAS 操作可以不停地检查某个需要满足的条件, 一旦条件满足就退出for循环.
	// 这就相当于,只要条件未被满足,当前的流程就会被一直"阻塞" 在这里
	for {
		if atomic.CompareAndSwapInt32(&num2, 10, 0) {
			fmt.Println("The second number has gone to zero.")
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func TestDemo3(t *testing.T) {
	var num int32 = 50

	go func() {
		for i := 0; i < 50; i++ {
			atomic.AddInt32(&num, 1)
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 50; i++ {
			atomic.StoreInt32(&num, 18)
			time.Sleep(time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
	fmt.Println(atomic.LoadInt32(&num))
}
