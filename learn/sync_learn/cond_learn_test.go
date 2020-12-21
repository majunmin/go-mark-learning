/**
  @Author: majm@ushareit.com
  @date: 2020/12/21
  @note:
**/
package sync_learn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	// 0: 表示有消息 1： 表示无消息
	var mailbox int8

	// lock变量的 Lock() 和 Unlock() 分别用于对其中写锁的锁定和解锁,它们与 sendCond变量 的含义是对应的
	var lock sync.RWMutex

	// sendCond是专门为放置情报而准备的条件变量，向信箱里放置情报,可以被视为`对共享资源的写操作`
	sendCond := sync.NewCond(&lock)
	// recvCond变量代表的是专门为获取情报而准备的条件变量,,可以看做是`对共享资源的读操作`
	recvCond := sync.NewCond(lock.RLocker())

	sign := make(chan struct{}, 1)

	max := 5

	// 发信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 1; i <= max; i++ {
			time.Sleep(1 * time.Second)
			lock.Lock()

			for mailbox == 1 { //信箱中有信 等待
				sendCond.Wait()
			}
			fmt.Printf("sender [%d]: the mailbox is empty. \n", i)
			mailbox = 1
			fmt.Printf("sender [%d]: the mailbox has message. \n", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)

	// 收信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 1; i <= max; i++ {
			time.Sleep(2 * time.Second)

			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			fmt.Printf("receiver [%d]: the mailbox is full. \n", i)
			mailbox = 0
			fmt.Printf("receiver [%d]: the letter has been received. \n", i)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
