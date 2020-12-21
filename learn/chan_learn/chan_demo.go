/**
  @Author: majm@ushareit.com
  @date: 2020/11/23
  @note:
**/
package chanchan

import (
	"fmt"
	"math/rand"
	"time"
)

var done = make(chan bool)
var msg string

// close(chan) 之后，其他从chan中取出来的值是类型零值
func AGoutinue() {
	msg = "hello.xlsx world!"
	//time.Sleep(time.Second * 2)
	done <- true
	//close(done)

}

func Operate() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	ch1 <- 4
	ch1 <- 5

	elem := <-ch1

	fmt.Printf("The first element of the chan1 is %v \n", elem)
}

//SelectOpe 只要有 default 语句 就不会被阻塞, select 就不会 其他case 里面的通道阻塞
func SelectOpe() {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index // 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。

	// 如果多个case 分支满足条件, 那么会随机选择一个执行
	//intChannels[0] <- 0
	//intChannels[1] <- 1
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}

func SelectForPend() {
	intChan := make(chan int, 1) // 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
