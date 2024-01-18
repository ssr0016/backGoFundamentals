package utils

// What is Panic?
// The idiomatic way of handling abnormal conditions in a Go program is using errors. Errors are sufficient for most of the abnormal conditions arising in the program.

// But there are some situations where the program cannot continue execution after an abnormal condition. In this case, we use panic to prematurely terminate the program. When a function encounters a panic, its execution is stopped, any deferred functions are executed and then the control returns to its caller. This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message, followed by the stack trace and then terminates. This concept will be more clear when we write an example program.

import (
	"fmt"
	"runtime/debug"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s", *firstName, *lastName)
	fmt.Printf("returned normally from fullName\n")
}

func PanicAndRecover() {
	firstName := "Samson"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

// Panics can also be caused by errors that happen during the runtime such as trying to access an index that is not present in a slice.

func slicePanic() {
	n := []int{1, 2, 3}
	fmt.Println(n[4])
	fmt.Println("returned normally from slicePanic")
}

func ImplementSlicePanic() {
	slicePanic()
	fmt.Println("returned normally from main")
}

// Defer Calls During a Panic

// Let’s recollect what panic does. When a function encounters a panic, its execution is stopped, any deferred functions are executed and then the control returns to its caller. This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message, followed by the stack trace and then terminates.

func fullName1(firstName *string, lastName *string) {
	defer fmt.Println("dferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s", *firstName, *lastName)
	fmt.Printf("returned normally from fullName\n")
}

func DeferCallsDuringAPanic() {
	defer fmt.Println("deferred call in main")
	firstName := "Samson"
	fullName1(&firstName, nil)
	fmt.Println("returned normally from main")
}

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func fullName3(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func RecoveringFromAPanic() {
	defer fmt.Println("deferred call in main")
	firstName := "Samson"
	fullName3(&firstName, nil)
	fmt.Println("returned normally from main")
}

func recoverInvalidAccess() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func invalidSliceAccess() {
	defer recoverInvalidAccess()
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normallyy returned from a")
}

func ImplementRecoverInvalidAccess() {
	invalidSliceAccess()
	fmt.Println("returned normally from main")
}

// Getting Stack Trace after Recover
// If we recover from a panic, we lose the stack trace about the panic. Even in the program above after recovery, we lost the stack trace.

func recoverFullName1() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
		debug.PrintStack()
	}
}

func fullName4(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func GettingStackTraceAfterRecover() {
	defer fmt.Println("deferred call in main")
	firstName := "Samson"
	fullName4(&firstName, nil)
	fmt.Println("returned normally from main")
}

// Panic, Recover and Goroutines

// Recover works only when it is called from the same goroutine which is panicking. It’s not possible to recover from a panic that has happened in a different goroutine.

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d +%d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	<-done

}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	done <- true
}

func PanicRecoverAndGoroutines() {
	sum(5, 0)
	fmt.Println("returned normally from main")
}
