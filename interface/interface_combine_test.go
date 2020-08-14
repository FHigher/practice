package inter

import "testing"

func TestCombine(t *testing.T) {
	var c i1
	c = &cat{}
	c.eat()

	var c2 i2
	c2 = &cat{}

	printi1(c2)
}
