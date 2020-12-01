/**
  @Author: majm@ushareit.com
  @date: 2020/11/25
  @note:
**/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	Client = &http.Client{}
)

func main() {
	// 测试自定义的 Get 函数
	GetData()

	// 当然，我们不是为了自定义 Get  Post 方法，才去实现 Do 方法的调用
	// 是因为有一些特殊的需求，比如要携带 cookie ， 连接超时时间设置等。
	MyConnection()

}

func MyConnection() {
	log.Println("--------------- 开始执行自定义的 Do 方法 ----------------")
	// 正常执行自定义的请求
	myDo()

	log.Println("--------------- 开始执行超时后的自定义的 Do 方法 ----------------")
	// 给 Client 设置一个超时，然后测试效果
	// 这里只是简单的验证超时，没有管 Transport 的设置
	Client.Timeout = 1e9
	myDo()
	// 可以看到会有一个链接超时的错误，如下面这样
	// Get http://localhost:8080: net/http: request canceled (Client.Timeout exceeded while awaiting headers)
}

func myDo() {
	// 先自定义一个 Request
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	ErrPrint(err)

	// 设置一个 csrf token, 服务器端会去验证这个
	req.Header.Set("_csrf", "12345")
	// 设置一个 Cookie, 服务器端会去验证这个
	req.Header.Set("Cookie", "name=majm")
	// 一会写完服务器端可以尝试分别将上面两行注释去测试

	resp, err := Client.Do(req)
	ErrPrint(err)
	defer resp.Body.Close()

	DataPrint(resp.Body)
	// 打印下状态码，看下效果
	fmt.Printf("返回的状态码是： %v\n", resp.StatusCode)
	fmt.Printf("返回的信息是： %v\n", resp.Body)
}

func GetData() {
	log.Println("---------custom get obtain data----------")
	url := "https://www.baidu.com"
	resp, err := MyGet(url)
	ErrPrint(err)
	defer resp.Body.Close()

	DataPrint(resp.Body)

}

func MyGet(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return Client.Do(request)
}

func DataPrint(body io.ReadCloser) {
	// 拿到数据
	bytes, err := ioutil.ReadAll(body)
	ErrPrint(err)

	// 这里要格式化再输出，因为 ReadAll 返回的是字节切片
	fmt.Printf("%s\n", bytes)
}

func ErrPrint(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
