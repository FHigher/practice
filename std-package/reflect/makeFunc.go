package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `k:"n"`
}

func (s Student) GetName() string {
	return s.Name
}

var caclAdd func(int, int) int64

var swap = func(args []reflect.Value) []reflect.Value {
	return []reflect.Value{args[1], args[0]}
}

func main() {
	var num = Student{"zhangsan"}
	//p := uint(10)
	//var num *uint = &p
	fmt.Println(reflect.TypeOf(num))
	//fmt.Println(reflect.ValueOf(num).Elem())

	switch reflect.TypeOf(num).Kind() {
	case reflect.Ptr:
		fmt.Println("Ptr type")
	default:
		fmt.Println("unknown")
	}

	fmt.Println(reflect.ValueOf(num))
	//fmt.Println(reflect.ValueOf(num).Elem().Field(0).Tag())
	fmt.Println(reflect.TypeOf(num).Field(0).Tag.Get("k"))
	fmt.Println(reflect.TypeOf(num).NumMethod())

	fmt.Println(reflect.TypeOf(new(error)).Elem())

	mutf := reflect.MakeFunc(reflect.ValueOf(&caclAdd).Elem().Type(), func(args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(args[0].Int() + args[1].Int())}
	})

	reflect.ValueOf(&caclAdd).Elem().Set(mutf)

	fmt.Println(mutf.Interface())
	fmt.Println(caclAdd(1, 2))

	makeSwap := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		fmt.Println(fn.Type())
		v := reflect.MakeFunc(fn.Type(), swap)
		fn.Set(v)
	}

	var intfn func(int, int) (int, int)
	makeSwap(&intfn)
	fmt.Println(intfn(2, 3))

	var floatfn func(float64, float64) (float64, float64)
	makeSwap(&floatfn)
	fmt.Println(floatfn(2.5, 3.6))

	ff := reflect.FuncOf([]reflect.Type{reflect.TypeOf(5)}, []reflect.Type{reflect.TypeOf("hello")}, false)
	fmt.Println(ff)
	fmt.Println(ff.Kind())
	fmt.Println(reflect.TypeOf(5))
}
