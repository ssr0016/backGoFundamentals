package utils

import (
	"fmt"
	"math/rand"
)

func Switch() {
	finger := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
}

// Duplicate cases are not allowed

func DuplicateCasesIsNotAllowed() {
	// finger := 4
	// fmt.Printf("Finger %d is ", finger)
	// switch finger {
	// case 1:
	// 	fmt.Println("Thumb")
	// case 2:
	// 	fmt.Println("Index")
	// case 3:
	// 	fmt.Println("Middle")
	// case 4:
	// 	fmt.Println("Ring")
	// case 4: //duplicate case
	// 	fmt.Println("Another Ring")
	// case 5:
	// 	fmt.Println("Pinky")

	// }
}

func DefaultCase() {
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: // default case
		fmt.Println("incorrect finger number")
	}
}

func MultipleExpressionInCase() {
	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u": // multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

}

func ExpressionlessSwitch() {
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 101", num)
	}
}

func number() int {
	num := 15 * 5
	return num
}

func Fallthrough() {
	switch num := number(); { // num is not used
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is less than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	}
}

func FallthroughHappensEvenWhenTheCaseEvaluatesToFalse() {
	switch num := 25; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
	}
}

func BreakingSwitch() {
	switch num := -5; {
	case num < 50:
		if num < 0 {
			break
		}
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is less than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	}
	// In the above program num is -5. When the control reaches the if statement in line no. 10, the condition is satisfied since num < 0.
	//  The break statement terminates the switch before it completes and the program doesn’t print anything.
}

func BreakingTheOuterForLoop() {
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("Generated even number %d", i)
			break randloop
		}
	}
	// Please note that if the break statement is used without the label, the switch statement will only be broken and the loop will continue running.
	//  So labeling the loop and using it in the break statement inside the switch is necessary to break the outer for loop.
}
