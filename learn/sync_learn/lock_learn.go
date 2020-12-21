/**
  @Author: majm@ushareit.com
  @date: 2020/12/19
  @note:
**/
package sync_learn

import "sync"

// sync.Mutex 值类型
var lck sync.Mutex

// 重复上锁 加锁
// ![](https://raw.githubusercontent.com/majunmin/image/master/img%E4%BA%92%E6%96%A5%E9%94%81%E7%9A%84%E9%87%8D%E5%A4%8D%E9%94%81%E5%AE%9A%E5%92%8C%E9%87%8D%E5%A4%8D%E8%A7%A3%E9%94%81.png)
func Demo1() {
	lck.Lock()
	// 重复上锁 会引发 panic: fatal error: all goroutines are asleep - deadlock!
	//lck.Lock()
	lck.Unlock()
	// 重复解锁   会引发panic: fatal error: sync: unlock of unlocked mutex
	lck.Unlock()
}
