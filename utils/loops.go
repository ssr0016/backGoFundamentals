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

func ForLoopContinue() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf(" %d", i)
	}
}

// Nested for loops
func ForLoopNested() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Labels
func ForLoopLabels() {
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}

func ForLoopLabelsBreak() {
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}
}

func ForLoopLabelsOuter() {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
}

func MoreExamples() {
	i := 0
	for i <= 10 {
		fmt.Printf("%d", i)
		i += 2
	}
}

func MoreExamplesSemicolonOmmitted() {
	i := 0
	for i <= 10 { // sermicolons are ommitted and only condition is present
		fmt.Printf("%d", i)
		i += 2
	}
}

func MultipleVariablesInForLoop() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { // multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}

// for {
// 	// infinite loop
// }

