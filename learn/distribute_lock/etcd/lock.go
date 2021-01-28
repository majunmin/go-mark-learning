/**
  @Author: majm@ushareit.com
  @date: 2021/1/28
  @note:
**/
package etcd

import (
	"github.com/zieckey/etcdsync"
	"log"
)

type Lock struct {
}

func etcdLockDemo() {
	// etcdsync使用的是 etcdv2 版本api
	mu, err := etcdsync.New("/lock", 10, []string{
		"http://127.0.0.1:55001",
		"http://127.0.0.1:55002",
		"http://127.0.0.1:55000",
	})
	if mu == nil || err != nil {
		log.Println("etcdsync.New failed")
		return
	}

	err = mu.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
		return
	}

	log.Println("etcd sync locked success")
	log.Println("DO SOMETHING")

	err = mu.Unlock()
	if err != nil {
		log.Println("etcd sync unlock failed")
	}
	log.Println("etcd sync unlock ok")
}
