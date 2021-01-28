/**
  @Author: majm@ushareit.com
  @date: 2021/1/28
  @note:
**/
package distribute_lock

import (
	"fmt"
	"sync"
)

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

//TryLock not block, return lock result
func (l Lock) TryLock() bool {
	var lockResult bool
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

//
func (l Lock) UnLock() {
	l.c <- struct{}{}
}

var countx int

func tryLockDemo() {
	var wg sync.WaitGroup
	lock := NewLock()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !lock.TryLock() {
				fmt.Println("lock failed")
				return
			}
			countx++
			fmt.Println("current counter", countx)
			lock.UnLock()
		}()
	}

	wg.Wait()
}
