package utils

import (
	"fmt"
	"math"
)

// A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type.

// func (t Type) methodName(parameter list) {
// }

//Sample methods

type Employee1 struct {
	name     string
	salary   int
	currency string
}

func (e Employee1) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func CallingDisplaySalaryMethods() {
	emp1 := Employee1{
		name:     "John",
		salary:   5000,
		currency: "$",
	}

	emp1.displaySalary() //Calling displaySalary() method of Employee type
}

// Methods vs Functions
func displaySalary(e Employee1) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func CallingDisplaySalaryFunction() {
	emp1 := Employee1{
		name:     "John",
		salary:   5000,
		currency: "$",
	}
	displaySalary(emp1)
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func CallingRectangleAndCicleTypeUsingMethod() {
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle is %d\n", r.Area())

	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle is %f\n", c.Area())
}

// Pointers Recievers vs Value Receivers
type Employee2 struct {
	name string
	age  int
}

// Method with value receiver
func (e Employee2) changeName(newName string) {
	e.name = newName
}

// Method with pointer receiver
func (e *Employee2) changeAge(newAge int) {
	e.age = newAge
}

func CallingPointersRecieversVsValueReceivers() {
	e := Employee2{
		name: "John",
		age:  25,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}

// Methods of anonymous struct fields

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address:%s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

func CallingMethodsOfAnonymousStructFields() {
	p := person{
		firstName: "John",
		lastName:  "Doe",
		address: address{
			city:  "San Francisco",
			state: "California",
		},
	}

	p.fullAddress() //Accessing fullAddress() method of address struct
	// fmt.Println("\n", p)
}

// Value recievers in methods vs Value arguments in functions

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", r.length*r.width)
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", r.length*r.width)
}

func CallingValueRecieversInMethodsVsValueArgumentsInFunctions() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)

	p := &r //calling value receiver with a pointer
	p.area()
}

// Pointer receivers in methods vs Pointer arguments in functions

type rectangle1 struct {
	length int
	width  int
}

func perimeter(r *rectangle1) {
	fmt.Printf("Perimeter Function output: %d\n", 2*(r.length+r.width))
}

func (r *rectangle1) perimeter() {
	fmt.Printf("Perimeter Method output: %d\n", 2*(r.length+r.width))
}

func CallingPointerRecieversInMethodsVsPointerArgumentsInFunctions() {
	r := rectangle1{
		length: 10,
		width:  5,
	}
	p := &r // pointer to r
	perimeter(p)
	p.perimeter()

	/*
		cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	//perimeter(r)

	r.perimeter() // calling pointer receiver with a value
}

// Methods with non-struct receivers

type MyInt int

func (a MyInt) add(b MyInt) MyInt {
	return a + b
}

func MethodsWithNonStructReceivers() {
	num1 := MyInt(5)
	num2 := MyInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
