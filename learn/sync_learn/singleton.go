/**
  @Author: majm@ushareit.com
  @date: 2020/11/22
  @note:
**/
package sync_learn

import (
	"sync"
	"sync/atomic"
)

var (
	instance    *Singleton
	initialized uint32
	mu          sync.Mutex
)

type Singleton struct{}

var instance2 *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() { instance2 = &Singleton{} })
	return instance2
}

func Instance() *Singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &Singleton{}
	}
	return instance
}
