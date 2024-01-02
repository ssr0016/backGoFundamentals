package utils

import (
	"fmt"
	"time"
)

// Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as lightweight threads. The cost of creating a Goroutine is tiny when compared to a thread. Hence it’s common for Go applications to have thousands of Goroutines running concurrently.

func hello1() {
	fmt.Println("Hello World!")
}

func Goroutine1() {
	go hello1()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func StartingMultipleGoroutines() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main terminated")
}
