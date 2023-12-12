package utils

import (
	"fmt"
)

func DeclaringAConstant() {
	const a = 50
	fmt.Println(a)
}

func DeclaringAGroupOfConstants() {
	const (
		name    = "samson"
		age     = 50
		country = "Philippines"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// var a = math.Sqrt(4)   // allowed
	// const b = math.Sqrt(4) // not allowed

	const n = "Sam"
	var names = n
	fmt.Printf("type %T value %v", names, names)

	var defaultName = "Sam"
	type myString string
	var customName myString = "Samson"
	customName = myString(defaultName)

	fmt.Println(customName)
}

// func BooleanConstants() {
// 	const trueConst = true
// 	type myBool bool
// 	var defaultBool = trueConst       // allowed
// 	var customBool myBool = trueConst // allowed
// 	defaultBool = (customBool)     // not allowed
// }

func NemeriConstants() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type is %T, f's type is %T, c's type is %T", i, f, c)

	const a1 = 5
	var intVar1 int = a
	var int32Var2 int32 = a
	var float64Var3 float64 = a
	var complex64Var1 complex64 = a
	fmt.Println("intVar1", intVar1, "\nint32Var2", int32Var2, "\nfloat64Var3", float64Var3, "\ncomplex64Var1", complex64Var1)

}

func NumericExpressions() {
	var a = 5.9 / 8
	fmt.Printf("a's type is %T and value is %v", a, a)

}
