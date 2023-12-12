package main

import "goFundamental/utils"

func main() {
	// Variables
	// utils.DeclaringASingleVariable()
	// utils.DeclaringAVariableWithAnInitialValue()
	// utils.MultipleVariableDeclaration()
	// utils.DeclareVariablesOfDifferentTypes()
	// utils.ShortHansDeclaration()
	// utils.VariablesAssignedValuesAndComputedDuringRuntime()

	//Types
	// utils.Bool()
	// utils.SignedIntegers()
	// utils.FloatingPointTypes()
	// utils.ComplessTypes()
	// utils.StringType()
	// utils.TypeConversion()

	//Constant
	// utils.DeclaringAConstant()
	// utils.DeclaringAGroupOfConstants()
	// utils.NemeriConstants()
	// 	utils.NumericExpressions()

	// Functions
	// utils.CalculateBill(100, 5)
	// price, no := 90, 6
	// totalPrice := utils.CalculateBill(price, no)
	// fmt.Println("Total price is", totalPrice)
	// utils.CalculateBill(price, no)

	//Multiple return values
	// area, perimeter := rectProps(10.8, 5.6)
	// fmt.Printf("Area %f Perimeter %f", area, perimeter)
	// area, perimeter := 10.8, 5.6
	// utils.RectProps(area, perimeter)

	//Named return values
	// area2, perimeter2 := utils.RectProps2(10.8, 5.6)
	// fmt.Printf("Area %f Perimeter %f", area2, perimeter2)
	// area2, perimeter2 := 10.8, 5.6
	// utils.RectProps2(area2, perimeter2)

	//Blank identifier
	// area, _ := utils.RectProps3(10.8, 5.6)
	// fmt.Printf("Area %f", area)
	area, _ := 10.8, 5.6
	utils.RectProps3(area, 5.6)

}