package utils

import (
	"fmt"
	"goFundamental/computer"
)

// A struct is a user-defined type that represents a collection of fields. It can be used in places where it makes sense to group the data into a single unit rather than having each of them as separate values.

// Declaring a struct
type Employee struct {
	firstname string
	lastname  string
	age       int
	salary    int
}

func DeclaringAStruct() {
	// creating struct specifying field names
	emp1 := Employee{
		firstname: "John",
		lastname:  "Doe",
		age:       25,
		salary:    5000,
	}

	// creating struct without specifying field names
	emp2 := Employee{"Thomas", "Paul", 30, 6000}

	fmt.Println("Employee 1:", emp1)
	fmt.Println("Employee 2:", emp2)
}

// It is possible to declare structs without creating a new data type. These types of structs are called anonymous structs.
func CreatingAnonymousStructs() {
	emp3 := struct {
		firstname string
		lastname  string
		age       int
		salary    int
	}{
		firstname: "John",
		lastname:  "Doe",
		age:       25,
		salary:    5000,
	}

	fmt.Println("Employee 3:", emp3)
}

func AccessingIndividualFieldsOfAStruct() {
	emp6 := Employee{
		firstname: "John",
		lastname:  "Doe",
		age:       25,
		salary:    5000,
	}

	fmt.Println("First Name:", emp6.firstname)
	fmt.Println("First Name:", emp6.lastname)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)
	emp6.salary = 6500
	fmt.Printf("Salary: $%d\n", emp6.salary)
}

func ZeroValueOfAStruct() {
	var emp4 Employee // Zero valued struct
	fmt.Println("First Name:", emp4.firstname)
	fmt.Println("Last Name:", emp4.lastname)
	fmt.Println("Age:", emp4.age)
	fmt.Println("Salary: ", emp4.salary)

	emp5 := Employee{
		firstname: "John",
		lastname:  "Paul",
	}
	fmt.Println("First Name:", emp5.firstname)
	fmt.Println("Last Name:", emp5.lastname)
	fmt.Println("Age:", emp5.age)
	fmt.Println("Salary: ", emp5.salary)
}

func PointersToAStruct() {
	emp8 := &Employee{
		firstname: "John",
		lastname:  "Doe",
		age:       25,
		salary:    5000,
	}
	fmt.Println("First Name:", (*emp8).firstname)
	fmt.Println("Age:", (*emp8).age)

	emp7 := &Employee{
		firstname: "John",
		lastname:  "Doe",
		age:       25,
		salary:    5000,
	}
	fmt.Println("First Name:", emp7.firstname)
	fmt.Println("Age:", emp7.age)
}

type Person struct {
	string
	int
}

func AnonymousFields() {
	p1 := Person{
		string: "Samson",
		int:    30,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}

type Person1 struct {
	name    string
	age     int
	address Address
}

type Address struct {
	city  string
	state string
}

func NestedStructs() {
	p := Person1{
		name: "Samson",
		age:  30,
		address: Address{
			city:  "Manila",
			state: "Philippines",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}

type Person2 struct {
	name string
	age  int
	Address
}

func PromotedFields() {
	p := Person2{
		name: "Samson",
		age:  30,
		Address: Address{
			city:  "Manila",
			state: "Philippines",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   // city is promoted field
	fmt.Println("State:", p.state) // state is promoted field
}

func ExportedStructsAndFields() {
	spec := computer.Spec{
		Maker: "Dell",
		Price: 5000,
	}
	fmt.Println("Maker", spec.Maker)
	fmt.Println("Price", spec.Price)

	spec1 := computer.Spec{
		Maker: "Lenovo",
		Price: 10000,
		Model: "ThinkPad",
	}
	fmt.Println("Maker", spec1.Maker)
	fmt.Println("Price", spec1.Price)

}

type name struct {
	firstname string
	lastname  string
}

func StructsEquality() {
	name1 := name{
		firstname: "Steve",
		lastname:  "Jobs",
	}
	name2 := name{
		firstname: "Steve",
		lastname:  "Jobs",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstname: "Steve",
		lastname:  "Jobs",
	}
	name4 := name{
		firstname: "Steve",
	}
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}

// In the program above image struct type contains a field data which is of type map. maps are not comparable, hence image1 and image2 cannot be compared. If you run this program, the compilation will fail with error
type image struct {
	data map[int]int
}

func StructvariablesAreNotComparableIfTheyContainFieldsThatAreNotComparable() {
	image1 := image{
		data: map[int]int{
			0: 155,
		}}
	image2 := image{
		data: map[int]int{
			0: 155,
		}}
	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	}
}
