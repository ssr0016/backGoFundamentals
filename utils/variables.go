package utils

import (
	"fmt"
	"math"
)

func DeclaringASingleVariable() {
	// Var name type is the syntax to declare a single variable
	var age int // variable declaration
	fmt.Println("My age is", age)
	age = 29 // assignment
	fmt.Println("My age is", age)
	age = 54 // assignment
	fmt.Println("My age is", age)
}

func DeclaringAVariableWithAnInitialValue() {
	// var name type = initiavalue
	var age int = 29 // variable declaration with initial value
	fmt.Println("My age is", age)

	// var name = initiavalue
	var age1 = 29 // type will ber inferred
	fmt.Println("My age is", age1)
}

func MultipleVariableDeclaration() {
	// var name1, name2 type = initialvalue1, initialvalue2 is the syntax for multiple variable declaration.

	var width, height int = 100, 50 // declaring multiple variables
	fmt.Println("width is", width, "height is", height)

	// The type can be removed if the variables have an initial value. Since the above program has initial values for variables, the int type can be removed.
	var width1, height1 = 100, 50 // "int" is dropped
	fmt.Println("width is", width1, "height is", height1)

	var width2, height2 int
	fmt.Println("width is", width2, "height is", height2)
	width2 = 100
	height2 = 50
	fmt.Println("width is", width2, "height is", height2)
}

func DeclareVariablesOfDifferentTypes() {
	// 	var (
	//       name1 = initialvalue1
	//       name2 = initialvalue2
	// 		)
	var (
		name   = "samson"
		age    = 30
		height int
	)

	fmt.Println("My name is", name, " age is", age, "and height is", height)

}

func ShortHansDeclaration() {
	count := 10
	fmt.Println("Count =", count)

	//multiple variables in a single line using short hand syntax.

	name, age := "samson", 30 // short hand declaration
	fmt.Println("My name is", name, " age is", age)

	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("change b is", b, "c is", c)

	// it will occur error because no new variables
	// a1, b1 := 20, 30 // a and b declared
	// fmt.Println("a1 is", a1, "b2 is", b1)
	// a1, b1 := 40, 50 // error, no new variables
}

func VariablesAssignedValuesAndComputedDuringRuntime() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)

	// age3 := 29      // age is int
	// age3 = "naveen" // error since we are trying to  assign a string to a variable of type int
}
