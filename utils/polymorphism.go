package utils

import "fmt"

// Polymorphism using an interface
// Any type which provides definition for all the methods of an interface is said to implicitly implement that interface. This will be more clear as we discuss an example of polymorphism shortly.

// A variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go.

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d\n", netincome)
}

func PolymorphismImplementation() {
	project1 := FixedBilling{projectName: "Project1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project2", biddedAmount: 100000}
	project3 := TimeAndMaterial{projectName: "Project3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
