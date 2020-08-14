package inter

import (
	"fmt"
	"testing"
)

func TestByteCount(t *testing.T) {
	var b Bytecounter
	b.Write([]byte("hello"))
	fmt.Println(b)
	b.Write([]byte("world"))
	fmt.Println(b)

	fmt.Fprintf(&b, "hello %s", "world")
	fmt.Println(b)
}

func TestStringer(t *testing.T) {
	person := &Person{"Peter", 25}
	fmt.Println(person)
}
