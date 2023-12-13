package utils

import "fmt"

// if statement syntax
// if condition {

// }

func EvenOdd() {
	num := 10
	if num%2 == 0 { // checks if number is even
		fmt.Println("The number", num, "is even")
		return
	}
	fmt.Println("The number", num, "is odd")
}

// if condition {
// } else {
// }
func EvenOdd1() {
	num1 := 11
	if num1%2 == 0 { // checks if number is even
		fmt.Println("The number", num1, "is even")
	} else {
		fmt.Println("The number", num1, "is odd")
	}
}

// If … else if … else statement
// if condition1 {
// ...
// } else if condition2 {
// ...
// } else {
// ...
// }

func LessOrGreater() {
	num3 := 99
	if num3 <= 50 {
		fmt.Println(num3, "is less than or equal to 50")
	} else if num3 >= 51 && num3 <= 100 {
		fmt.Println(num3, "is between 51 and 100")
	} else {
		fmt.Println(num3, "is greater than 100")
	}
}

// If with assignment
// if assignment-statement; condition {
// }

func EvenOdd2() {
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

func IdiomaticEvenOdd() {
	if num4 := 10; num4%2 == 0 {
		fmt.Println(num4, "is even")
	} else {
		fmt.Println(num4, "is odd")
	}
}

// The idiomatic way of writing the above program in Go’s philosophy is to avoid the else and return from the if if the condition is true.
func ExampleIdiomaticEvenOdd() {
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")
}
