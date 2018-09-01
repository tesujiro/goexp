package main

import (
	"fmt"
	"reflect"
)

func main() {
	//makeFuncAndCallSample3()
	makeVariadicFuncAndCallSample1()
}

func makeFuncAndCallSample1() {
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

func makeFuncAndCallSample2() {
	//
	// FUNCTION: addOne(int) int
	//

	// first parameter of MakeFunc : function interface
	inType := make([]reflect.Type, 1)
	outType := make([]reflect.Type, 1)
	inType[0] = reflect.TypeOf(1)                      // 1: int
	outType[0] = reflect.TypeOf(1)                     // 1: int
	funcType := reflect.FuncOf(inType, outType, false) // false: not variadic

	// second parameter of MakeFunc
	addOneVal := func(in []reflect.Value) []reflect.Value {
		if in[0].Type().Kind() == reflect.Int {
			return []reflect.Value{reflect.ValueOf(in[0].Interface().(int) + 1)}
		}
		return []reflect.Value{}
	}

	addOneFuncVal := reflect.MakeFunc(funcType, addOneVal)

	// Call Func : addOne(2)
	args := []reflect.Value{reflect.ValueOf(2)}
	result := addOneFuncVal.Call(args)
	fmt.Println(result[0].Int())
}

func makeFuncAndCallSample3() {
	//
	// FUNCTION: plus(interface{},interface{}) interface{}
	//

	// first parameter of MakeFunc : function interface
	inType := make([]reflect.Type, 2)
	outType := make([]reflect.Type, 1)
	reflectValueType := reflect.TypeOf(reflect.Value{})
	inType[0] = reflectValueType                       // ival: interface{}
	inType[1] = reflectValueType                       // ival: interface{}
	outType[0] = reflectValueType                      // ival: interface{}
	funcType := reflect.FuncOf(inType, outType, false) // false: not variadic

	// second parameter of MakeFunc
	plusVal := func(in []reflect.Value) []reflect.Value {
		val_kind := func(v reflect.Value) reflect.Kind {
			return v.Interface().(reflect.Value).Type().Kind()
		}
		val_interface := func(v reflect.Value) interface{} {
			return v.Interface().(reflect.Value).Interface()
		}

		if val_kind(in[0]) == reflect.Int && val_kind(in[1]) == reflect.Int {
			ret := val_interface(in[0]).(int) + val_interface(in[1]).(int)
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf(ret))}
		}
		if val_kind(in[0]) == reflect.String && val_kind(in[1]) == reflect.String {
			ret := val_interface(in[0]).(string) + val_interface(in[1]).(string)
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf(ret))}
		}
		return []reflect.Value{reflect.ValueOf(reflect.ValueOf(nil))}
	}

	PlusFuncVal := reflect.MakeFunc(funcType, plusVal)

	// Call
	var args, result []reflect.Value
	val_val := func(v interface{}) reflect.Value {
		return reflect.ValueOf(reflect.ValueOf(v))
	}

	// Call Func : int + int
	args = []reflect.Value{val_val(2), val_val(2)}
	result = PlusFuncVal.Call(args)
	fmt.Println(result[0].Interface())

	// Call Func : string + string
	args = []reflect.Value{val_val("abc"), val_val("def")}
	result = PlusFuncVal.Call(args)
	fmt.Println(result[0].Interface())
}

func makeVariadicFuncAndCallSample1() {
	// first parameter of MakeFunc
	var add func(...int) int

	// second parameter of MakeFunc
	addVal := func(in []reflect.Value) []reflect.Value {
		if in[0].Kind() == reflect.Slice {
			result := 0
			for i := 0; i < in[0].Len(); i++ {
				result += in[0].Index(i).Interface().(int)
			}
			return []reflect.Value{reflect.ValueOf(result)}
		}
		return []reflect.Value{reflect.ValueOf(0)}
	}

	makeFunc := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), addVal)
		fn.Set(v)
	}
	makeFunc(&add)

	// At first, call the func directly
	fmt.Println(add(1, 2, 3))

	// Convert func to reflect.Value
	addFuncVal := reflect.ValueOf(add)

	// Call Func
	args := []reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3), reflect.ValueOf(4)}
	result := addFuncVal.Call(args)
	fmt.Println(result[0].Int())
}
