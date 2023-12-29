package utils

import "fmt"

// In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface. It is much similar to the OOP world. Interface specifies what methods a type should have and the type decides how to implement these methods.

// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func DeclaringAndImplementingAnInterface() {
	name := MyString("Samson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", v.FindVowels())
}

// =======================================================================

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

// salary of permanent employee is the sum of basicpay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// salary of contract employee is the basicpay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpenses(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expenses Per Month $%d\n", expense)
}

func PracticalUseOfAnInterface() {
	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp3 := Contract{
		empId:    3,
		basicpay: 3000,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp3}
	totalExpenses(employees)
}

type SalaryCalculator1 interface {
	CalculateSalary1() int
}

type Permanent1 struct {
	empId    int
	basicpay int
	pf       int
}

type Contract1 struct {
	empId    int
	basicpay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

// salary of permanent1 employee is the sum of basicpay and pf
func (p1 Permanent1) CalculateSalary1() int {
	return p1.basicpay + p1.pf
}

// salary of contract1 employee is the basicpay alone
func (c1 Contract1) CalculateSalary1() int {
	return c1.basicpay
}

// salary of freelancer
func (f Freelancer) CalculateSalary1() int {
	return f.ratePerHour * f.totalHours
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/

func totalExpenses1(s1 []SalaryCalculator1) {
	expense := 0
	for _, v := range s1 {
		expense = expense + v.CalculateSalary1()
	}
	fmt.Printf("Total expenses Per Month $%d\n", expense)
}

func PracticalUseOfAnInterface1() {
	pemp1 := Permanent1{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent1{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp3 := Contract1{
		empId:    3,
		basicpay: 3000,
	}
	freelancer1 := Freelancer{
		empId:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	freelancer2 := Freelancer{
		empId:       5,
		ratePerHour: 100,
		totalHours:  100,
	}
	employees := []SalaryCalculator1{pemp1, pemp2, cemp3, freelancer1, freelancer2}
	totalExpenses1(employees)
}

// =====================================

type Worker interface {
	Work()
}

type Person5 struct {
	name string
}

func (p Person5) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func InterfaceInternalRepresentation() {
	p := Person5{
		name: "John",
	}
	var w Worker = p
	describe(w)
	w.Work()
}

// =====================================

func describe1(i interface{}) {
	fmt.Printf("Type: %T Value: %v\n", i, i)
}

func EmptyInterface() {
	s1 := "Hello World"
	describe1(s1)

	i := 55
	describe1(i)

	strt := struct {
		name string
	}{
		name: "John",
	}
	describe1(strt)
}

// =====================================
func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

func TypeAssertion() {
	var s interface{} = 56
	assert(s)
}

func assert1(i1 interface{}) {
	v, ok := i1.(int)
	fmt.Println(v, ok)
}

func TypeAssertion1() {
	var s1 interface{} = 56
	assert1(s1)
	var i1 interface{} = "JL"
	assert1(i1)
}

// ==================================
// A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements. It is similar to switch case. The only difference being the cases specify types and not values as in normal switch.

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("unknown type\n")
	}
}

func TypeSwitch() {
	findType("Hello")
	findType(5)
	findType(5.5)
}

type Describer interface {
	Describe()
}

type Person6 struct {
	name string
	age  int
}

func (p Person6) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func findType1(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func TypeSwitch1() {
	findType1("Hello")
	p := Person6{
		name: "John",
		age:  30,
	}
	findType1(p)
}
