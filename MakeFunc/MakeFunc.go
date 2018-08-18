package main

import (
	"fmt"
	"reflect"
)

func main() {
	makeFuncAndCallSample()
}

func makeFuncAndCallSample() {
	// first parameter of MakeFunc
	var addOne func(int) int

	// second parameter of MakeFunc
	addOneVal := func(in []reflect.Value) []reflect.Value {
		if in[0].Type().Kind() == reflect.Int {
			return []reflect.Value{reflect.ValueOf(in[0].Interface().(int) + 1)}
		}
		return []reflect.Value{}
	}

	makeFunc := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), addOneVal)
		fn.Set(v)
	}

	makeFunc(&addOne)

	// At first, call the func directly
	fmt.Println(addOne(1))

	// Convert func to reflect.Value
	addOneFuncVal := reflect.ValueOf(addOne)

	// Call Func
	args := []reflect.Value{reflect.ValueOf(2)}
	result := addOneFuncVal.Call(args)
	fmt.Println(result[0].Int())
}
