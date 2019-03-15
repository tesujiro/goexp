package main

import (
	"fmt"
	"math/big"
)

func main() {
	var f32a, f32b float32
	f32a, f32b = 1.2, 1.5
	fmt.Printf("f32a=%v\n", f32a)
	fmt.Printf("f32b=%v\n", f32b)
	fmt.Printf("f32a*f32b=%v\n", f32a*f32b)
	fmt.Printf("f32b-f32a=%v\n", f32b-f32a)

	var f64a, f64b float64
	f64a, f64b = 1.2, 1.5
	fmt.Printf("f64a=%v\n", f64a)
	fmt.Printf("f64b=%v\n", f64b)
	fmt.Printf("f64b-f64a=%v\n", f64b-f64a)

	var bigFloata, bigFloatb big.Float
	bigFloata.SetFloat64(1.2)
	bigFloatb.SetFloat64(1.5)
	fmt.Printf("bigFloata=%.10g\n", &bigFloata)
	fmt.Printf("bigFloatb=%.10g\n", &bigFloatb)
	var bigF big.Float
	bigF.SetPrec(32)
	fmt.Printf("bigFloatb*bigFloata=%.10g\n", bigF.Mul(&bigFloatb, &bigFloata))
	fmt.Printf("bigFloatb-bigFloata=%.10g\n", bigF.Sub(&bigFloatb, &bigFloata))
}
