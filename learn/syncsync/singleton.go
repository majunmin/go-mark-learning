/**
  @Author: majm@ushareit.com
  @date: 2020/11/22
  @note:
**/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	instance    *Singleton
	initialized uint32
	mu          sync.Mutex
)

type Singleton struct{}

// 将标准代码提取出来 就成了 标准库里的 sync.Once
func main() {
	// custom:
	instance1 := Instance()
	fmt.Println(&instance1)

	println(GetInstance())

}

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
