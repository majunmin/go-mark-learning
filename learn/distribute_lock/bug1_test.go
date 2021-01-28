/**
  @Author: majm@ushareit.com
  @date: 2021/1/28
  @note:
**/
package distribute_lock

import "testing"

func TestCondurrent(t *testing.T) {
	concurrentModify()
	// locked concurrentModify
	lockConcurrentModify()

	tryLockDemo()
}
