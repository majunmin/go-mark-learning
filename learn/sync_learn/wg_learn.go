/**
  @Author: majm@ushareit.com
  @date: 2020/12/22
  @note:
**/
package sync_learn

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func coordinateWithChan() {
	sign := make(chan struct{}, 2)

	var num int32 = 0
	fmt.Printf("The number: %d [with chan struct{}]\n", num)
	max := int32(10)

	for i := 0; i < 4; i++ {
		go addNum(&num, 1, max, func() {
			sign <- struct{}{}
		})
	}

	<-sign
	<-sign
	<-sign
	<-sign
	fmt.Println(num)
}

func coornateWithWg() {
	num := int32(0)
	var wg sync.WaitGroup
	wg.Add(4)

	for i := 0; i < 4; i++ {
		go addNum(&num, int32(i), 10, wg.Done)
	}

	wg.Wait()
	fmt.Println(num)
}

// addNum 用于原子地增加numP所指的变量的值。
func addNum(numP *int32, id, max int32, deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(numP)
		if curNum >= max {
			break
		}
		newNum := curNum + 2
		if atomic.CompareAndSwapInt32(numP, curNum, newNum) {
			fmt.Printf("The num: %d [%d - %d] \n", newNum, id, i)
		} else {
			fmt.Printf("atomic CAS operte failed! [%d - %d] \n", id, i)
		}
	}

}
