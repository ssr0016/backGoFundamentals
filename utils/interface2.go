package utils

import "fmt"

type DescriberII interface {
	DescribeII()
}

type PersonInsterfaceII struct {
	name string
	age  int
}

func (p PersonInsterfaceII) DescribeII() { // implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type AddressInteraceII struct {
	state string
	city  string
}

// methods with pointer receivers
func (a *AddressInteraceII) DescribeII() { // implemented using pointer receiver
	fmt.Printf("%s, %s\n", a.state, a.city)
}

func ImplementingInterfacesUsingPointerReceiversVsValueReceivers() {
	var d1 DescriberII
	p1 := PersonInsterfaceII{
		"samson",
		30,
	}
	d1 = p1
	d1.DescribeII()
	p2 := PersonInsterfaceII{
		"jl",
		30,
	}
	d1 = &p2
	d1.DescribeII()

	var d2 DescriberII
	a := AddressInteraceII{
		"Manila",
		"Philippines",
	}
	d2 = &a
	d2.DescribeII()
	a1 := AddressInteraceII{
		"Manila",
		"Philippines",
	}
	d2 = &a1
	d2.DescribeII()

}

// ==============================================================================

type SalaryCalulatorII interface {
	DisplaySalaryII()
}

type LeaveCalculatorII interface {
	CalculateLeavesLeftII() int
}

type EmployeeII struct {
	firtName    string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e EmployeeII) DisplaySalaryII() {
	fmt.Printf("%s %s has salary %d\n", e.firtName, e.lastName, (e.basicPay + e.pf))
}

func (e EmployeeII) CalculateLeavesLeftII() int {
	return e.totalLeaves - e.leavesTaken
}

func ImplementingMultipleInterfaces() {
	e := EmployeeII{
		firtName:    "Sams",
		lastName:    "Recs",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 10,
	}

	var s SalaryCalulatorII = e
	s.DisplaySalaryII()
	var l LeaveCalculatorII = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeftII())
}

// =================================================

//Embedded interfaces

type SalaryCalculatorEmbeded interface {
	DisplaySalaryEmbeded()
}

type LeaverCalculatorEmbeded interface {
	CalculateLeavesLeftEmbeded() int
}

type EmployeeOperations interface {
	SalaryCalculatorEmbeded
	LeaverCalculatorEmbeded
}

type EmployeeEmbeded struct {
	firtName    string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e EmployeeEmbeded) DisplaySalaryEmbeded() {
	fmt.Printf("%s %s has salary %d\n", e.firtName, e.lastName, (e.basicPay + e.pf))
}

func (e EmployeeEmbeded) CalculateLeavesLeftEmbeded() int {
	return e.totalLeaves - e.leavesTaken
}

func EmbeddingInterfaces() {
	e := EmployeeEmbeded{
		firtName:    "Sams",
		lastName:    "Recs",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 10,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalaryEmbeded()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeftEmbeded())
}

// =================================================
// Zero Value of Interface

type DescriberZeroValueOfInterface interface {
	DescribeZeroValueOfInterface()
}

func ZeroValueOfInterface() {
	var d1ZeroValueOfInterface DescriberZeroValueOfInterface
	if d1ZeroValueOfInterface == nil {
		fmt.Printf("d1ZeroValueOfInterface is nil and has type %T value %v\n", d1ZeroValueOfInterface, d1ZeroValueOfInterface)
	}
}
