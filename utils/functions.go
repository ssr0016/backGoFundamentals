package utils

import (
	"fmt"
)

// The syntax for declaring a function
// func functionname(parametername type) returntype{

// }

// FunctionsDeclaration

// func calculateBill(price int, no int) int {
// 	var totalPrice = price * no
// 	return totalPrice
// }

func CalculateBill(price, no int) int {
	var totalPrice = price * no
	fmt.Println("Total price is", totalPrice)
	return totalPrice
}

// Multiple return values
func RectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	fmt.Println("Area is", area, "Perimeter is", perimeter)
	return area, perimeter
}

// Named return values
func RectProps2(length, width float64) (area2, perimeter2 float64) {
	area2 = length * width
	perimeter2 = (length + width) * 2
	fmt.Println("Area is", area2, "Perimeter is", perimeter2)
	return // no explicit return values
}

// Blank identifier
func RectProps3(length, width float64) (float64, float64) {
	var area3 = length * width
	var perimeter3 = (length + width) * 2
	fmt.Println("Area is", area3, "Perimeter is", perimeter3)
	return area3, perimeter3
}
