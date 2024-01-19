package utils

import "fmt"

// What are first class functions?
// A language that supports first class functions allows functions to be assigned to variables, passed as arguments to other functions and returned from other functions.

type add func(a int, b int) int

func UserDefinedFunctionTypes() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}

// Higher-order functions

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func PassingFunctionsAsArgumentsToOtherFunctions() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}

func simple1() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func ReturningFunctionsFromOtherFunctions() {
	s := simple1()
	fmt.Println(s(60, 7))
}

func Closures() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func Closures1() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func Closures2() {
	s1 := student{
		firstName: "John",
		lastName:  "Doe",
		grade:     "A",
		country:   "USA",
	}
	s2 := student{
		firstName: "Jane",
		lastName:  "Doe",
		grade:     "B",
		country:   "USA",
	}
	s3 := student{
		firstName: "Jack",
		lastName:  "Doe",
		grade:     "C",
		country:   "USA",
	}
	s := []student{s1, s2, s3}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func Closures3() {
	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 2
	})
	fmt.Println(r)
}
