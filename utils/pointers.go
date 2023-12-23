package utils

import "fmt"

// What is a pointer?
// A pointer is a variable that stores the memory address of another variable.

func DeclaringPointers() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("Address of b is", a)
}

func ZeroValueOfApointer() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
}

func CreatingPointersUsingTheNewFunction() {
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)

}

func DereferencingAPointer1() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
}
func DereferencingAPointer2() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)
}

func change(val *int) {
	*val = 25
}

func PassingPointerToAFunction() {
	a := 58
	fmt.Println("Value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("Value of a after function call is", a)
}

func hello() *int {
	i := 5
	return &i
}

func ReturningPointerFromAFunctions() {
	d := hello()
	fmt.Println("Value of d", *d)
}

//Note

func modify1(arr *[3]int) {
	(*arr)[0] = 90
}
func modify2(arr *[3]int) {
	(arr)[0] = 90
}

func modify3(sls []int) {
	sls[0] = 90
}

func DoNotPassAPointerToAnArrayAsAnArgumentToAFunction1() {
	a := [3]int{89, 90, 91}
	modify1(&a)
	fmt.Println(a)
}
func DoNotPassAPointerToAnArrayAsAnArgumentToAFunction2() {
	a := [3]int{89, 90, 91}
	modify2(&a)
	fmt.Println(a)
}

func DoNotPassAPointerToAnArrayAsAnArgumentToAFunction3() { // it is the best approach do not pass a pointer in array as argument
	a := [3]int{89, 90, 91}
	modify3(a[:])
	fmt.Println(a)
}

// Go does not support pointer arithmetic
// Go does not support pointer arithmetic which is present in other languages like C and C++.

// func GoDoesNotSupportPointerArithmetic() {
// 	b := [...]int{109, 110, 111}
// 	p := &b
// 	p++
// }
