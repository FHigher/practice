package inter

import "fmt"

type i1 interface {
	eat()
}

type i2 interface {
	i1
	walk()
}

type cat struct{}

func (*cat) eat() {
	fmt.Println("cat eat")
}

func (*cat) walk() {
	fmt.Println("cat walk")
}

func printi1(i i1) {
	i.eat()
}

func printi2(i i2) {
	i.eat()
	i.walk()
}
