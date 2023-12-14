package utils

import "fmt"

// for loop syntax

// for initialisation; condition; post {

func ForLoop() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
}

func ForLoopBreak() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break // loop is terminated if i > 5
		}
		fmt.Printf(" %d", i)
	}
	fmt.Printf("\nline after for loop")
}
