package main

import (
	"fmt"
)

func main() {
	b := []byte{1, 4}
	fmt.Printf("%b\n", int16(b[0])<<8)
	fmt.Printf("%b\n", int16(b[1]))
	fmt.Printf("%b\n", 260)

	fmt.Println(Tint16(b))

	a := make([]byte, 2)
	PutInt16(a, 260)
	fmt.Println(a)
}

func Tint16(b []byte) int16 {
	return int16(b[1]) | int16(b[0])<<8
}

func PutInt16(b []byte, v int16) {
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}
