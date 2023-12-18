package utils

import "fmt"

// Array

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

func printArrayMultidimentionalArrays(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func MultidimentionalArrays() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printArrayMultidimentionalArrays(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArrayMultidimentionalArrays(b)

}

// Slices

// Declaration & Creating a slice

func Slice() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // creates a slice from a[1] to a [3]
	fmt.Println(b)
	// output [77 78 79]
}

func SliceAnotherExampleAndScenario1() {
	c := []int{1, 2, 3} // creates and array and returns a slice reference
	fmt.Println(c)
}

func ModifyingASlice() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}

func SliceAnotherExampleAndScenario2() {
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
}

func LenghtAndCapacityOfASlice() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice is %d and capacity is %d\n", len(fruitslice), cap(fruitslice)) ///length of fruitslice is 2 and capacity is 6
}

func CreatingASliceUsingMake() {
	i := make([]int, 5, 5)
	fmt.Println(i)
}

func AppendingToASlice1() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled
}

func AppendingToASlice2() {
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "john", "sebastian", "vinay")
		fmt.Println("names contents:", names)
	}
}
func AppendingToASlice3() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}

type slice struct {
	Length        int
	Capacity      int
	ZerothElement *byte
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func PassingASliceToAFunction() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                                //function modifies the slice
	fmt.Println("slice before function call", nos) //modifications are visible outside
}

func MultidimentionalSlices() {
	pls := [][]string{
		{"C", "C++"},
		{"Javascript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func countries() []string {
	countries := []string{"USA", "SINGAPORE", "INDIA", "GERMANY", "FRANCE"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func MemoryOptimisation() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
