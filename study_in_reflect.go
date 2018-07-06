package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//basic1()
	//basic2()
	//basic3()
	//basic4()
	//callFunc()
	makeFunc1()
	makeFunc2()
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
	fmt.Println("\tType():", reflect.ValueOf(i).Type())
	fmt.Println("\tKind():", reflect.ValueOf(i).Kind())
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

func basic4() {
	rvNil := reflect.Value{}
	fmt.Println("reflect.Value{}")
	fmt.Println("\tType:", reflect.TypeOf(rvNil))
	fmt.Println("\tValue:", reflect.ValueOf(rvNil))
}

func callFunc() {
	fn := fmt.Println
	fmt.Println("fmt.Println")
	fmt.Println("\tType:", reflect.TypeOf(fn))
	fmt.Println("\tValue:", reflect.ValueOf(fn))
	fmt.Println("\tKind():", reflect.TypeOf(fn).Kind())
	fmt.Println("\tNumIn():", reflect.TypeOf(fn).NumIn())
	fmt.Println("\tIn(0):", reflect.TypeOf(fn).In(0))
	fmt.Println("\tIsVariadic():", reflect.TypeOf(fn).IsVariadic())

	fnValue := reflect.ValueOf(fn)

	arg1 := reflect.ValueOf("Hello world!")
	fnValue.Call([]reflect.Value{arg1})
}

func makeFunc1() {
	// Function 1 : func(int) int
	// see below comments
	double := func(in []reflect.Value) []reflect.Value {
		i := in[0].Interface().(int)

		return []reflect.Value{reflect.ValueOf(i * 2)}
	}

	inType := []reflect.Type{reflect.TypeOf(1)}
	outType := []reflect.Type{reflect.TypeOf(1)}
	funcType := reflect.FuncOf(inType, outType, false)

	// func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
	// ==> args must be []reflect.Value
	// ==> return values of func must be []reflect.Value
	fn := reflect.MakeFunc(funcType, double)

	arg1 := reflect.ValueOf(123)
	ret := fn.Call([]reflect.Value{arg1})
	fmt.Println("func double(int) int")
	fmt.Println("\tdouble(123)=", ret[0])
}

func makeFunc2() {
	// Function 2 : func(int) (int,error)
	double := func(in []reflect.Value) []reflect.Value {
		var errValue reflect.Value
		var errorType = reflect.ValueOf([]error{nil}).Index(0).Type()
		var reflectValueErrorNilValue = reflect.ValueOf(reflect.New(errorType).Elem())
		errValue = reflectValueErrorNilValue
		if in[0].Kind() != reflect.Int {
			errValue = reflect.ValueOf(errors.New("1st parameter not int"))
		}
		i := in[0].Interface().(int)

		return []reflect.Value{reflect.ValueOf(reflect.ValueOf(i * 2)), reflect.ValueOf(errValue)}
	}

	inType := []reflect.Type{reflect.TypeOf(1)}
	outType := []reflect.Type{reflect.TypeOf(reflect.Value{}), reflect.TypeOf(reflect.Value{})}
	funcType := reflect.FuncOf(inType, outType, false)

	fn := reflect.MakeFunc(funcType, double)

	arg1 := reflect.ValueOf(123)
	ret := fn.Call([]reflect.Value{arg1})
	fmt.Println("func double(int) (int,error)")
	fmt.Println("\tdouble(123)=", ret[0].Interface(), ret[1].Interface())
}

func makeFunc7() {
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
