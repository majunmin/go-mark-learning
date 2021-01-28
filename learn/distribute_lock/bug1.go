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

var count int
var lCount int
var mu sync.Mutex

func concurrentModify() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func lockConcurrentModify() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			lCount++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(lCount)
}
