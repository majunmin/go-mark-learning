# io

io包 中 `io.Reader`接口 以及` io.Writer`接口 的 **扩展接口** 以及 **实现类型** 以及他们的作用



1. io.Reader

- `Read(p []byte) (n int, err error)`

2. io.Writer

- `Write(p []byte) (n int, err error)`

3. io.Closer

- `Close() error`

4. io.Seeker

- `Seek(offset int64, whence int) (int64, error)`

5. io.ReadWriter
6. io.ReadCloser
7. io.WriteCloser
8. io.ReadSeeker
9. io.WriteSeeker
10. io.LimitedReader
11. io.SectionReader
12. io.teeReader
13. io.ReadWriteCloser
14. io.ReadWriteSeeker

## io.Copy()

```go
// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the number of bytes
// copied and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
//
// If src implements the WriterTo interface,
// the copy is implemented by calling src.WriteTo(dst).
// Otherwise, if dst implements the ReaderFrom interface,
// the copy is implemented by calling dst.ReadFrom(src).
func Copy(dst Writer, src Reader) (written int64, err error) {
    return copyBuffer(dst, src, nil)
}

// CopyN copies n bytes (or until an error) from src to dst.
// It returns the number of bytes copied and the earliest
// error encountered while copying.
// On return, written == n if and only if err == nil.
//
// If dst implements the ReaderFrom interface,
// the copy is implemented using it.
func CopyN(dst Writer, src Reader, n int64) (written int64, err error) {
    written, err = Copy(dst, LimitReader(src, n))
    if written == n {
        return n, nil
    }
    if written < n && err == nil {
        // src stopped early; must have been EOF.
        err = EOF
    }
    return
}


// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
// @param buf: 缓冲区长度, default buffer size is 32 * 1024
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
    // If the reader has a WriteTo method, use it to do the copy.
    // Avoids an allocation and a copy.
    if wt, ok := src.(WriterTo); ok {
        return wt.WriteTo(dst)
    }
    // Similarly, if the writer has a ReadFrom method, use it to do the copy.
    if rt, ok := dst.(ReaderFrom); ok {
        return rt.ReadFrom(src)
    }
    if buf == nil {
        size := 32 * 1024
        if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
            if l.N < 1 {
                size = 1
            } else {
                size = int(l.N)
            }
        }
        buf = make([]byte, size)
    }
    for {
        nr, er := src.Read(buf)
        if nr > 0 {
            nw, ew := dst.Write(buf[0:nr])
            if nw > 0 {
                written += int64(nw)
            }
            if ew != nil {
                err = ew
                break
            }
            if nr != nw {
                err = ErrShortWrite
                break
            }
        }
        if er != nil {
            if er != EOF {
                err = er
            }
            break
        }
    }
    return written, err
}

```

## io.ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)

> `ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)`:从输入流中读取至少min个字节到buf中，直到把buf读取满或者遇到 error，包括EOF。

1. # 如果 n < min, 肯定返回错误
2. 没有数据可读返回n=0和EOF
3. n < min ,并且遇到EOF,返回n和`ErrUnexpectedEOF`
4. 如果参数 min>len(buf) ,返回0,`ErrShortBuffer`
5. 如果读取了至少min个字节，即使遇到错误也会返回err=nil

```go
// ReadAtLeast reads from r into buf until it has read at least min bytes.
// It returns the number of bytes copied and an error if fewer bytes were read.
// The error is EOF only if no bytes were read.
// If an EOF happens after reading fewer than min bytes,
// ReadAtLeast returns ErrUnexpectedEOF.
// If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.
// On return, n >= min if and only if err == nil.
// If r returns an error having read at least min bytes, the error is dropped.
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) {
	if len(buf) < min {
		return 0, ErrShortBuffer
	}
	for n < min && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn
	}
	if n >= min {
		err = nil
	} else if n > 0 && err == EOF {
		err = ErrUnexpectedEOF
	}
	return
}
```
