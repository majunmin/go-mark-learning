/**
  @Author: majm@ushareit.com
  @date: 2020/12/26
  @note:
**/
package sync_learn

import (
	"fmt"
	"sync"
	"time"
)

var pool sync.Pool

type Person struct {
	Name string
}

func initPool() {
	pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new person")
			return new(Person)
		},
	}
}

func poolDemo1() {
	initPool()
	person := pool.Get().(*Person)

	fmt.Println("get object from pool : ", person)

	person.Name = "first"

	pool.Put(person)

	go func() {
		fmt.Println("get object from pool : ", pool.Get().(*Person))
	}()

	go func() {
		fmt.Println("get object from pool : ", pool.Get().(*Person))
	}()

	fmt.Println("get object from pool : ", pool.Get().(*Person))
	fmt.Println("get object from pool : ", pool.Get().(*Person))
	fmt.Println("get object from pool : ", pool.Get().(*Person))

	time.Sleep(time.Second)

}
