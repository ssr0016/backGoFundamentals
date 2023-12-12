package utils

import (
	"fmt"
	"unsafe"
)

// The following are the basic types available in Go

// bool
// Numeric Types
// 	int8, int16, int32, int64, int
// 	uint8, uint16, uint32, uint64, uint
// 	float32, float64
// 	complex64, complex128
// 	byte
// 	rune
// string

func Bool() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
}

func SignedIntegers() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is, size of b is %d", unsafe.Sizeof(b))
}

func UnsignedIntegers() {

}

func FloatingPointTypes() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum is", sum, "diff is", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

func ComplessTypes() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}

// Other numeric types
// byte is an alias of uint8
// rune is an alias of int32

func StringType() {
	first := "Samson"
	last := "Recluta"
	name := first + " " + last
	fmt.Println("My name is", name)
}

func TypeConversion() {
	i := 55   // int
	j := 55.5 // float64
	// sum := i + j     // int + float64
	// fmt.Println(sum) // (mismatched types int and float64)

	sum := i + int(j)
	fmt.Println(sum)

	i1 := 10
	var j1 float64 = float64(i1) // this statement will not work without explicit conversion
	fmt.Println("j1", j1)
}
