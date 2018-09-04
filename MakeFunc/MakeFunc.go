package main

import (
	"fmt"
	"reflect"
)

func main() {
	//makeFuncAndCallSample1()
	//makeFuncAndCallSample2()
	//makeFuncAndCallSample3()
	//makeVariadicFuncAndCallSample1()
	//makeVariadicFuncAndCallSample2()
	//makeVariadicFuncAndCallSample3()
	//callGoFuncByReflect1()
	//callGoFuncByReflect2()
	//callGoFuncByReflect3()
	makeFuncWithPtrArgsAndCallSample1()
}

func makeFuncAndCallSample1() {
	// first parameter of MakeFunc
	var addOne func(int) int

	// second parameter of MakeFunc
	addOneVal := func(in []reflect.Value) []reflect.Value {
		if in[0].Type().Kind() != reflect.Int {
			return []reflect.Value{} //ERROR
		}
		return []reflect.Value{reflect.ValueOf(in[0].Interface().(int) + 1)}
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
		if in[0].Type().Kind() != reflect.Int {
			return []reflect.Value{}
		}
		return []reflect.Value{reflect.ValueOf(in[0].Interface().(int) + 1)}
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
		if in[0].Kind() != reflect.Slice || in[0].Index(0).Type().Kind() != reflect.Int {
			return []reflect.Value{reflect.ValueOf(0)} // 0: ERROR
		}
		result := 0
		for i := 0; i < in[0].Len(); i++ {
			result += in[0].Index(i).Interface().(int)
		}
		return []reflect.Value{reflect.ValueOf(result)}
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

func makeVariadicFuncAndCallSample2() {
	//
	// FUNCTION: add(...int) int
	//

	// first parameter of MakeFunc : function interface
	inType := make([]reflect.Type, 1)
	outType := make([]reflect.Type, 1)
	inType[0] = reflect.TypeOf([]int{})               // arg1: int slice
	outType[0] = reflect.TypeOf(1)                    // 1: int
	funcType := reflect.FuncOf(inType, outType, true) // true: variadic

	// second parameter of MakeFunc
	addVal := func(in []reflect.Value) []reflect.Value {
		if in[0].Kind() != reflect.Slice || in[0].Index(0).Type().Kind() != reflect.Int {
			return []reflect.Value{reflect.ValueOf(0)} // 0: ERROR
		}
		result := 0
		for i := 0; i < in[0].Len(); i++ {
			result += in[0].Index(i).Interface().(int)
		}
		return []reflect.Value{reflect.ValueOf(result)}
	}

	addFuncVal := reflect.MakeFunc(funcType, addVal)

	// Call Func : add(10,11,12)
	args := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(11), reflect.ValueOf(12)}
	result := addFuncVal.Call(args)
	fmt.Println(result[0].Int())
}

func makeVariadicFuncAndCallSample3() {
	//
	// FUNCTION: add(...interface{}) interface{}
	//

	// first parameter of MakeFunc : function interface
	inType := make([]reflect.Type, 1)
	outType := make([]reflect.Type, 1)
	reflectValueType := reflect.TypeOf(reflect.Value{})
	inType[0] = reflect.TypeOf([]reflect.Value{})     // []interface{}
	outType[0] = reflectValueType                     // interface{}
	funcType := reflect.FuncOf(inType, outType, true) // true: variadic

	// second parameter of MakeFunc
	addVal := func(in []reflect.Value) []reflect.Value {
		if in[0].Kind() != reflect.Slice {
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf("not defined as variadic"))}
		}

		switch in[0].Index(0).Interface().(reflect.Value).Type().Kind() {
		case reflect.Int:
			result_int := 0
			for i := 0; i < in[0].Len(); i++ {
				result_int += in[0].Index(i).Interface().(reflect.Value).Interface().(int)
			}
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf(result_int))}
		case reflect.String:
			result_str := ""
			for i := 0; i < in[0].Len(); i++ {
				result_str += in[0].Index(i).Interface().(reflect.Value).Interface().(string)
			}
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf(result_str))}
		default:
			return []reflect.Value{reflect.ValueOf(reflect.ValueOf("invalid value"))}
		}
	}

	AddFuncVal := reflect.MakeFunc(funcType, addVal)

	// Call
	var args, result []reflect.Value
	val_val := func(v interface{}) reflect.Value {
		return reflect.ValueOf(reflect.ValueOf(v))
	}

	// Call Func : int + int + int
	args = []reflect.Value{val_val(101), val_val(102), val_val(103)}
	result = AddFuncVal.Call(args)
	fmt.Println(result[0].Interface())

	// Call Func : string + string + string
	args = []reflect.Value{val_val("ABC"), val_val("DEF"), val_val("GHI")}
	result = AddFuncVal.Call(args)
	fmt.Println(result[0].Interface())
}

func callGoFuncByReflect1() {
	FuncVal := reflect.ValueOf(fmt.Println)

	// Call
	var args, result []reflect.Value
	val_val := func(v interface{}) reflect.Value {
		return reflect.ValueOf(reflect.ValueOf(v))
	}

	// Call Func : string + string + string
	args = []reflect.Value{val_val("ABC"), val_val("DEF"), val_val("GHI")}
	result = FuncVal.Call(args)
	fmt.Println(result[0].Interface())
}

func callGoFuncByReflect2() {
	FuncVal := reflect.ValueOf(fmt.Printf)

	// Call
	var args, result []reflect.Value
	val_val := func(v interface{}) reflect.Value {
		return reflect.ValueOf(reflect.ValueOf(v))
	}

	// Call Func : string + string + string
	args = []reflect.Value{reflect.ValueOf("%v:%v\n"), val_val("DEF"), val_val("GHI")}
	result = FuncVal.Call(args)
	fmt.Println(result[0].Interface())
}

func callGoFuncByReflect3() {
	addInt := func(args ...int) int {
		var result int
		for _, v := range args {
			result += v
		}
		return result
	}
	FuncVal := reflect.ValueOf(addInt)
	// Call
	var args, result []reflect.Value
	args = []reflect.Value{reflect.ValueOf(123), reflect.ValueOf(456)}
	result = FuncVal.Call(args)
	fmt.Println(result[0].Interface())
}

func makeFuncWithPtrArgsAndCallSample1() {
	// first parameter of MakeFunc
	var swap func(*int, *int)

	// second parameter of MakeFunc
	swapVal := func(in []reflect.Value) []reflect.Value {
		if in[0].Type().Kind() != reflect.Ptr {
			return []reflect.Value{} //ERROR
		}
		fmt.Printf("0:%#v\n", in[0].Elem())
		fmt.Printf("1:%#v\n", in[1].Elem())
		fmt.Printf("0:type %v\n", in[0].Elem().Type())
		fmt.Printf("1:type %v\n", in[1].Elem().Type())
		v0 := in[0].Elem()
		v1 := in[1].Elem()
		in[0].Elem().Set(v1)
		in[1].Elem().Set(v0)
		//return []reflect.Value{reflect.ValueOf(in[0].Interface().(int) + 1)}
		return nil
	}

	makeFunc := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), swapVal)
		fn.Set(v)
	}
	makeFunc(&swap)

	// At first, call the func directly
	a, b := 1, 2
	fmt.Println("a:", a, " b:", b)
	swap(&a, &b)
	fmt.Println("a:", a, " b:", b)

	/*
		// Convert func to reflect.Value
		addOneFuncVal := reflect.ValueOf(addOne)

		// Call Func
		args := []reflect.Value{reflect.ValueOf(2)}
		result := addOneFuncVal.Call(args)
		fmt.Println(result[0].Int())
	*/
}
