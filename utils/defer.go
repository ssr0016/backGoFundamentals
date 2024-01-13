package utils

import (
	"fmt"
	"sync"
)

// What is Defer?
// Defer statement is used to execute a function call just before the surrounding function where the defer statement is present returns. The definition might seem complex but it’s pretty simple to understand by means of an example.

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func Defer() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	largest(nums)
}

// ==========================================

type personDeferredStruct struct {
	firstName string
	lastName  string
}

func (p personDeferredStruct) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func DeferredMethods() {
	p := personDeferredStruct{
		firstName: "John",
		lastName:  "Doe",
	}
	defer p.fullName()
	fmt.Printf("Welcome\n")
}

// ==========================================

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func ArgumentsEvaluation() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
}

// =============================================
// When a function has multiple defer calls, they are pushed on to a stack and executed in Last In First Out (LIFO) order.

// We will write a small program which prints a string in reverse using a stack of defers.

func StackOfDefers() {
	name := "Samson"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range name {
		defer fmt.Printf("%c", v)
	}
}

// ==============================================
// Defer is used in places where a function call should be executed irrespective of the code flow. Let’s understand this with the example of a program which makes use of WaitGroup. We will first write the program without using defer and then we will modify it to use defer and understand how useful defer is.

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v has area %d\n", r, area)
}

func PracticalUserOfDefer() {
	var wg sync.WaitGroup
	r1 := rect{67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
