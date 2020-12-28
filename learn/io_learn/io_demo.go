/**
  @Author: majm@ushareit.com
  @date: 2020/12/28
  @note:
**/
package io_learn

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func ioDemo() {
	comment := "Because these interfaces and primitives wrap lower-level operations with various implementations, " +
		"unless otherwise informed clients should not assume they are safe for parallel execution."

	basicReader := strings.NewReader(comment)
	basicWriter := new(strings.Builder)

	lReader := io.LimitReader(basicReader, 10)
	_ = interface{}(lReader).(io.Reader)

	secReader := io.NewSectionReader(basicReader, 98, 89)
	_ = interface{}(secReader).(io.Reader)
	_ = interface{}(secReader).(io.ReaderAt)
	_ = interface{}(secReader).(io.Seeker)

	teeReader := io.TeeReader(basicReader, basicWriter)
	_ = interface{}(teeReader).(io.Reader)

	multiWriter := io.MultiWriter(basicWriter)
	_ = interface{}(multiWriter).(io.Writer)

	pipeReader, pipeWriter := io.Pipe()
	_ = interface{}(pipeReader).(io.Reader)
	_ = interface{}(pipeReader).(io.Closer)
	_ = interface{}(pipeWriter).(io.Writer)
	_ = interface{}(pipeWriter).(io.Closer)
}

func copyNDemo() {
	content := `
CopyN copies n bytes (or until an error) from src to dst.
It returns the number of bytes copied and the earliest
error encountered while copying.
On return, written == n if and only if err == nil.

If dst implements the ReaderFrom interface,
the copy is implemented using it.`
	reader := strings.NewReader(content)
	dst := new(strings.Builder)

	n, err := io.CopyN(dst, reader, 58)

	if err != nil {
		fmt.Printf("copy error, err= [%s]", err.Error())
	} else {
		fmt.Printf("Written (%d): %q \n", n, dst.String())
	}

	n, err = io.CopyN(dst, reader, 58)
	if err != nil {
		fmt.Printf("copy error, err= [%s]", err.Error())
	} else {
		fmt.Printf("Written (%d): %q \n", n, dst.String())
	}
}

func limitReaderDemo1() {
	content := `LimitReader returns a Reader that reads from r
 but stops with EOF after n bytes.
 The underlying implementation is a *LimitedReader.`
	reader := strings.NewReader(content)
	limitReader := io.LimitReader(reader, 10)
	buff := make([]byte, 7)
	n, err := limitReader.Read(buff)
	for i := 0; i < 5; i++ {
		executeIfNoErr(err, func() {
			fmt.Printf("Read(%d): %q\n", n, buff[:n])
		})
	}
}

// limitReader 的使用场景
// limitReader提供了一个保护措施
func limitReaderDemo() {
	// 建立连接
	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	// 发送请求, http 1.0 协议
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// 读取response
	var sb strings.Builder
	buf := make([]byte, 256)
	rr := io.LimitReader(conn, 102400)
	for {
		n, err := io.ReadAtLeast(rr, buf, 256)
		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				fmt.Println("read error:", err)
			}
			break
		}
		sb.Write(buf[:n])
	}
	// 显示结果
	fmt.Println("response:", sb.String())
	fmt.Println("total response size:", sb.Len())
}

//
func secReaderDemo() {
	content := "NewSectionReader returns a SectionReader that reads from r\n starting at offset off and stops with EOF after n bytes."
	reader := strings.NewReader(content)
	secReader := io.NewSectionReader(reader, 30, 30)
	for i := 0; i < 5; i++ {
		buff := make([]byte, 20)
		n, err := secReader.Read(buff)
		executeIfNoErr(err, func() {
			fmt.Printf("Read(%d): %q\n", n, buff[:n])
		})
	}
}

func executeIfNoErr(err error, f func()) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	f()
}
