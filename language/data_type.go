package main

import (
    "errors"
    "fmt"
    "strconv"
    "unicode"
)

func main() {
    //stringTest()
    //arrTest()
    //sliceTest()
    //mapTest()
    //structTest()
    interfaceTest()
    //errorTest()
}

/**
interface
 */
type Human interface {
    Say(word string) string
    Walk(s string) string
}

type Man struct {
    Name string
    Age int
}



func (m Man) Say(word string) string {
    return fmt.Sprintf("%s , my name is %s", word, m.Name)
}

func (m Man) Walk(s string) string {
    return fmt.Sprintf("Age is " + strconv.Itoa(m.Age) + ", and " + s)
}

func interfaceTest() {
    var m Man
    m.Name = "majm"
    m.Age= 26
    fmt.Println(m.Say("mike"))
    fmt.Println(m.Walk("go work!"))

    jack:= Man{Name: "jack", Age: 18}
    fmt.Println(jack.Say("majm"))
    fmt.Println(jack.Walk("gun"))
}

/**

 */
func errorTest() {
    var e error

    // 使用 errors定制错误信息
    e = errors.New("This is a test error")
    fmt.Println(e.Error())

    // 使用 fmt.Errorf 定义错误信息
    err:= fmt.Errorf("This is another test error")
    fmt.Println(err)

}

type Person struct {
    name string
    age int
}

/**
structTest
 */
func structTest() {
    var p Person
    p.name= "mike"
    p.age= 20
    fmt.Println(p)

    person := Person{name: "majm", age: 18}
    fmt.Println(person)

    p2:= new(Person) //new函数分配一个*指针*,指向 person 类型数据
    p2.age = 17
    p2.name = "zhangsanfeng"
    fmt.Println(*p2)
}


/*
test for map
 */
func mapTest() {
    var m1 map[int]string

    // 声明并初始化，初始化使用{} 或 make 函数(创建类型并分配空间)
    var m2 = make(map[string]string)
    m1 = map[int]string{} // 初始化 空 map

    m1[0] = "first"
    m1[1] = "second"
    m1[2] = "three"
    m2[`first`] = "一"
    m2[`second`] = "二"
    m2[`three`] = "三"

}

func sliceTest() {
    fmt.Println("===========slice test=============")
    var sl []int
    sl = append(sl, 1, 2, 3) // 往切片中追加值
    fmt.Println(sl)

    var arr = [5]int{1, 2, 3, 4, 5}
    fmt.Println(cap(arr)) // 5
    var sl1 = arr[1:2]    // 冒号:(左闭右开); 不填则默认为头|尾
    var sl2 = arr[:2]
    var sl3 = arr[1:]
    fmt.Println(sl1)
    fmt.Println(sl2)
    fmt.Println(sl3)

    sl1 = append(sl1, 1, 2, 3, 4, 5, 6, 7, 8)
    fmt.Println(sl1)
    fmt.Println(len(sl1))
    fmt.Println(cap(sl1)) // cap(slice) = len(slice) + 1

    // make create slice
    //sl1 := make([]int, 5) // 定义元素个数 = 5 的切片
    sl1 = make([]int, 5, 10) // 定义元素个数 = 5， 预留空间=10 的切片
    fmt.Println(len(sl1))
    fmt.Println(cap(sl1))
    strl := []string{"a", "b", "cc"}
    for i, s := range strl {
        fmt.Printf("index: %d, s = %v \n", i, s)
    }

}

/**
数组类型 arr
*/
func arrTest() {

    var arr1 [5]int
    arr1[0] = 1
    arr1[1] = 3
    arr1[2] = 2
    arr1[3] = 4
    fmt.Println(arr1)

}

func stringTest() {
    str := "pokyo hot"
    bytes := []byte(str)
    bytes[0] = 't'
    fmt.Printf("replace str(%v) to : %v \n", str, string(bytes))

    content := "西京热"
    runes := []rune(content)
    runes[0] = '东'
    fmt.Printf("replace content(%v) to : %v \n", content, string(runes))

    // count chinese word
    text := "hello沙河小丸子"
    count := 0
    for _, r := range text {
        if unicode.Is(unicode.Han, r) {
            count++
        }
    }
    fmt.Println(count)
}
