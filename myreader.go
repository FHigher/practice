package main

import (
	"errors"
	"fmt"
	"io"
)

var (
	// ErrReadClosed close
	ErrReadClosed = errors.New("输入流已关闭")
	// ErrReadToEmptySpace empty
	ErrReadToEmptySpace = errors.New("输入空间不能为nil")
)

// IMyReader reader interface
type IMyReader interface {
	io.Reader
	Close()
}

// MyReader 可读流
type MyReader struct {
	Data byte
	End  bool
}

func (r *MyReader) Read(b []byte) (int, error) {
	if r.End {
		return 0, ErrReadClosed
	}
	if b == nil {
		return 0, ErrReadToEmptySpace
	}
	length := len(b)
	for i := 0; i < length; i++ {
		b[i] = r.Data
	}

	return len(b), nil
}

// Close close
func (r *MyReader) Close() {
	r.End = true
}

func main() {
	var r IMyReader
	r = &MyReader{'A', false}

	//b := make([]byte, 8)
	var b []byte
	var counter int
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(n, err, b)
		fmt.Println(b[:n])
		counter++
		if counter >= 5 {
			r.Close()
		}
	}
}
