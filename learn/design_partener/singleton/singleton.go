/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package singleton

import "sync"

var once sync.Once

type singleton struct {
}

var instance *singleton

// 获取 单例对象
// 线程安全的单例模式
func GetSingleton() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
