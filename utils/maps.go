package utils

import "fmt"

//How to create map?
// make(map[keyType]valueType)
// employeeSalary := make(map[string]int)

// The above line of code creates a map named employeeSalary which has string keys and int values.

func CreateAMap() {
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)
}

// The syntax for adding new items to a map is the same as that of arrays. The program below adds some new employees to the employeeSalary map.
func AddingItemsToAMap() {
	employeeSalary := make(map[string]int)
	employeeSalary["John"] = 5000
	employeeSalary["Sarah"] = 6000
	employeeSalary[" Joerge"] = 6000
	fmt.Println("employeeSalary map contents", employeeSalary)
}

func InitializeAMapDuringDeclarationItself() {
	employeeSalary := map[string]int{
		"Batosai": 5000,
		"Samurai": 6000,
	}
	employeeSalary[" Joerge"] = 6000
	fmt.Println("employeeSalary map contents", employeeSalary)
}

// The zero value of a map is nil. If you try to add elements to a nil map, a run-time panic will occur. Hence the map has to be initialized before adding elements.
func ZeroValueOfAMap() {
	var employeeSalary map[string]int
	employeeSalary["Batosai"] = 5000
}

func RetrievingValueForAKeyFromAMap() {
	employeeSalary := map[string]int{
		"Batosai": 5000,
		"Samurai": 6000,
		"Joerge":  6000,
	}
	employee := "Batosai"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
}

func IfElementIsNotPresentINAMap() {
	employeeSalary := map[string]int{
		"Batosai": 5000,
		"Samurai": 6000,
	}
	fmt.Println("Salary of Sun is", employeeSalary["Sun"])
}

func CheckingIfAKeyExists() {
	employeeSalary := map[string]int{
		"Batosai": 5000,
		"Samurai": 6000,
	}
	newEmp := "Joerge"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")
}

func IterateOverAllElementsInAMap() {
	employeeSalary := map[string]int{
		"Batosai": 5000,
		"Samurai": 6000,
		"Joerge":  6000,
	}
	fmt.Println("Contents of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s]= %d\n", key, value)
	}
}

func DeletingItemsFromAMap() {
	employeeSalary := map[string]int{
		"Batosai": 5000,
		"Samurai": 6000,
		"Joerge":  6000,
	}
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "Joerge")
	fmt.Println("map after deletion", employeeSalary)
}

// ---------------------------------------------------------

type employee struct {
	salary  int
	country string
}

func MapOfStructs() {
	emp1 := employee{
		salary:  5000,
		country: "Philippines",
	}
	emp2 := employee{
		salary:  6000,
		country: "USA",
	}
	emp3 := employee{
		salary:  7000,
		country: "India",
	}
	employeeInfo := map[string]employee{
		"emp1": emp1,
		"emp2": emp2,
		"emp3": emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s, salary: %d, country: %s\n", name, info.salary, info.country)
	}
}

func LengthOfTheMap() {
	employeeSalary := map[string]int{
		"Batosai": 5000,
		"Samurai": 6000,
	}
	fmt.Println("length is", len(employeeSalary))
}

func MapsAreReferenceTypes() {
	employeeSalary := map[string]int{
		"Batosai": 5000,
		"Samurai": 6000,
		"Joerge":  6000,
	}
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["Joerge"] = 7000
	fmt.Println("Employee salary changed", employeeSalary)

}

// Maps can’t be compared using the == operator. The == can be only used to check if a map is nil.
func MapsEquality() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1
	if map1 == map2 {

	}
}
