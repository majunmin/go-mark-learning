/**
  @Author: majm@ushareit.com
  @date: 2020/11/21
  @note:
**/
package main

func main() {
	file, _ := OpenFile("./a.txt")
	var data []byte
	file.Read(17, data)

}

type File struct {
	Path string
}

// 函数
// OpenFile 类似一个构造函数， 构建 file对象
// CloseFile ReadFile 类似 File对象的成员方法
func OpenFile(name string) (f *File, err error) {
	// ...
	return nil, nil
}

func CloseFile(f *File) (err error) {
	// ...
	return nil
}

func ReadFile(f *File, offset int64, data []byte) int {
	// ...
	return 0
}

// 转化 File类型的 方法, 而不是 File对象的方法
// 方法的定义和 struct结构的定义必须在同一个`包`中

func (f *File) Read(offset int64, data []byte) int {
	return ReadFile(f, offset, data)
}

func (f *File) Close() error {
	return CloseFile(f)
}
