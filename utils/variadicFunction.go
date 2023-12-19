package utils

import "fmt"

// Functions in general accept only a fixed number of arguments. A variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by ellipsis …, then the function can accept any number of arguments for that parameter.

// Syntax
func Hello(a int, b ...int) {

}

// In the above piece of code, in line no. 1 we call hello with one argument 2 for the parameter b and we pass four arguments 6, 7, 8, 9 to b in the next line.

// Hello(1, 2)          // passing one argument "2" to b
// Hello(5, 6, 7, 8, 9) // passing arguments "6, 7 , 8 and 9" to b

// Example and understanding how variadic functions work
func Find1(num int, nums ...int) {
	fmt.Printf("Type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Printf("\n")
}

// Slice arguments vs Variadic arguments
func Find2(num int, nums ...int) {
	fmt.Printf("Type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Printf("\n")
}

// utils.Find2(89, []int{89, 90, 95})
// utils.Find2(45, []int{56, 67, 45, 90, 109})
// utils.Find2(78, []int{38, 56, 98})
// utils.Find2(87, []int{})

// Append is a variadic function
// func append(slice []Type, elems ...Type) []Type
func PassingASliceToAVariadicFunction(num int, nums ...int) {
	fmt.Printf("Type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Printf("\n")
}

func ImplementationPassingASliceToAVariadicFunction() {
	nums := []int{89, 90, 95}
	PassingASliceToAVariadicFunction(89, nums...)
}

func change1(s ...string) {
	s[0] = "Go"
}

func Gotcha1() {
	welcome := []string{"Hello", "World"}
	change1(welcome...)
	fmt.Println(welcome)
}

func change2(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func Gotcha2() {
	welcome := []string{"Hello", "World"}
	change2(welcome...)
	fmt.Println(welcome)
}
