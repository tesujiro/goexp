package main

import (
	"fmt"
	"reflect"
)

type xx struct {
	I int
	S string
}

func main() {
	fmt.Println("Hello, playground")

	XX := &xx{I: 123, S: "abc"}
	fmt.Println("XX:", XX)

	xv := reflect.ValueOf(XX).Elem()
	xt := reflect.TypeOf(XX).Elem()
	//fmt.Println("Kind:", xv.Kind().String())

	for i := 0; i < xv.NumField(); i++ {
		fmt.Println(i, xt.Field(i).Name, xv.Field(i))
	}

	num := 1234
	fmt.Println("Set Field 0 to ", num)
	xv.Field(0).Set(reflect.ValueOf(num))

	new := "new"
	fmt.Println("Set Field S to ", new)
	field := xv.FieldByName("S")
	fmt.Println("S=", field)
	field.SetString(new)

	for i := 0; i < xv.NumField(); i++ {
		fmt.Println(i, xt.Field(i).Name, xv.Field(i))
	}

}
