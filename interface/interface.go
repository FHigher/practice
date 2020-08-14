package inter

import (
	"fmt"
)

// Bytecounter byte counter
type Bytecounter int

func (b *Bytecounter) Write(str []byte) (int, error) {
	*b += Bytecounter(len(str))
	return len(str), nil
}

// Person person
type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("Person's Name is %s, Age is %d", p.Name, p.Age)
}
