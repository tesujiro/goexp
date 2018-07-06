package main

import (
	"fmt"
	"reflect"
)

func main() {
	//basic1()
	//basic2()
	//basic3()
	makeFunc()
}

func basic1() {
	i := 1
	fmt.Println("i int")
	fmt.Println("\tType:", reflect.TypeOf(i))
	fmt.Println("\tValue:", reflect.ValueOf(i))

	var f float64 = 1.1
	fmt.Println("f float64")
	fmt.Println("\tType:", reflect.TypeOf(f))
	fmt.Println("\tValue:", reflect.ValueOf(f))

	s := "Hello!"
	fmt.Println("s string")
	fmt.Println("\tType:", reflect.TypeOf(s))
	fmt.Println("\tValue:", reflect.ValueOf(s))
}

func basic2() {
	var k *int
	fmt.Println("k *int")
	fmt.Println("\tType:", reflect.TypeOf(k))
	fmt.Println("\tValue:", reflect.ValueOf(k))

	i := 1
	fmt.Println("&i")
	fmt.Println("\tType:", reflect.TypeOf(&i))
	fmt.Println("\tValue:", reflect.ValueOf(&i))

	fmt.Println("TypeOf(i)")
	fmt.Println("\tType:", reflect.TypeOf(reflect.TypeOf(i)))
	fmt.Println("\tValue:", reflect.ValueOf(reflect.TypeOf(i)))

	fmt.Println("ValueOf(i)")
	fmt.Println("\tType:", reflect.TypeOf(reflect.ValueOf(i)))
	fmt.Println("\tValue:", reflect.TypeOf(reflect.ValueOf(i)))
}

func basic3() {
	var a interface{}
	fmt.Println("a interface{}")
	fmt.Println("\tType:", reflect.TypeOf(a))
	fmt.Println("\tValue:", reflect.ValueOf(a))

	a = 123
	fmt.Println("a interface{} = 123")
	fmt.Println("\tType:", reflect.TypeOf(a))
	fmt.Println("\tValue:", reflect.ValueOf(a))
}

func makeFunc() {
	// swap is the implementation passed to MakeFunc.
	/*
		double := func(rv []reflect.Value) []reflect.Value {
			if rv[0].Kind() != reflect.Int {
				panic("func param panic!!")
			}
			i := rv[0].Interface().(int)

			return []reflect.Value{reflect.ValueOf(i * 2)}
		}
	*/
	double := func(i int) int {
		return i * 2
	}

	//inType := []reflect.Type{reflect.TypeOf(1)}
	inType := reflect.TypeOf(1)
	outType := inType
	funcType := reflect.FuncOf(inType, outType, false)

	reffn := reflect.MakeFunc(funcType, double)

	arg1 := reflect.ValueOf(123)
	ret := reffn.Call([]reflect.Value{arg1})
	fmt.Println(ret)
}

func makeFunc2() {
	// swap is the implementation passed to MakeFunc.
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}
	// makeSwap expects fptr to be a pointer to a nil function.
	makeSwap := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()     // fptr is a pointer to a function
		v := reflect.MakeFunc(fn.Type(), swap) // Make a function of the right type
		fn.Set(v)                              // Assign it to the value fn represents
	}
	// Make and call a swap function for ints.
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))
	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
}
