package utils

import "fmt"

// Declaration

func ArrayDeclaration1() {
	var a [3]int // int array with legnth 3
	a[0] = 12    // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

}

func ArrayDeclaration2() {
	a := [3]int{12, 78, 50} // shorty hand declaration to create array
	fmt.Println(a)
}

func ArrayDeclaration3() {
	a := [3]int{12}
	fmt.Println(a)
}

func ArrayDeclaration4() {
	a := [...]int{12, 78, 50} // ...makes the compiler determine the length
	fmt.Println(a)
}

func ArrayDeclaration5() {
	// a := [3]int{5, 78, 8}
	// var b [5]int
	// b = a // not possible since [3]int and [5]ing are distinct types
}

func ArrayAreValueTypes() {
	a := [...]string{"USA", "CHINA", "INDIA", "GERMANY", "FRANCE"}
	b := a // a copy of a is assigned to b
	b[0] = "SINGAPORE"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("Inside changeLocal", num)
}

func ImplementChangeLocalExample() {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) // num is passed by value
	fmt.Println("after passing to function ", num)
}

func LengthOfAnArray() {
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(a))
}

func IteratingArraysUsingRange1() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}
func IteratingArraysUsingRange2() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { //range returns both the index and the value
		fmt.Printf("%d th element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a is", sum)

	// for _, v := range a { //ignores index
}
