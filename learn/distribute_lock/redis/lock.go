/**
  @Author: majm@ushareit.com
  @date: 2021/1/28
  @note:
**/
package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var lockKey = "counter_lock"
var counterKey = "counter"

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	// lock
	resp := client.SetNX("lockKey", 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println("lock  failed!")
		return
	}

	getResp := client.Get(counterKey)
	cntVal, err := getResp.Int64()
	if err != nil || errors.Is(err, redis.Nil) {
		cntVal++
		setResp := client.Set(counterKey, cntVal, 0)
		_, err := setResp.Result()
		if err != nil {
			// log err
			fmt.Println("set value error!")
		}
	}

	fmt.Println("current counter is : ", cntVal)
	delResp := client.Del(counterKey)
	unlockSuccess, err := delResp.Result()
	if err != nil || unlockSuccess == 0 {
		fmt.Println("unlock failed!, err=", err)
	} else {
		fmt.Println("unlock success!")
	}

}

func redisDistributeLockDemo() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
