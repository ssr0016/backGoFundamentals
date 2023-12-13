package utils

import (
	"fmt"
	// "log"
)

// learnpackage
// -- go.mod
//   -- main.go
//      -- simpleinterest
//			-- simpleinterest.go

// package simpleinterest

// Calculate calculates and returns the simple interest for a principal p, rate of interest r and time t
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}

// Importing custom package
// import "learnpackage/simpleinterest"

func ImplementationExample() {
	fmt.Println("Simple interest calculation")
	p := 5000.0
	r := 10.0
	t := 1.0
	// si := simpleinterest.Calculate(p, r, t) // there is an error because tje simpleinterest is not imported
	si := Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}

// func init() {
// 	fmt.Println("Simple interest package initialized")
// }

// import (
// "fmt"
// "learnpackage/simpleinterest" //importing custom package
// log
// )

var p1, r1, t1 = 5000.0, 10.0, 1.0

// // init function to check if p, r and t are greater than zero
// func init() {
// 	println("Main package initialized")
// 	if p1 < 0 {
// 		log.Fatal("Principal is less than zero")
// 	}
// 	if r1 < 0 {
// 		log.Fatal("Rate of interest is less than zero")
// 	}
// 	if t1 < 0 {
// 		log.Fatal("Duration is less than zero")
// 	}
// }

func Main1() {
	fmt.Println("Simple interest calculation")
	si1 := Calculate(p1, r1, t1)
	fmt.Println("Simple interest is", si1)
}
