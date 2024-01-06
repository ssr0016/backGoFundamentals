package main

import "goFundamental/utils"

func main() {
	//---------------------------------------------------
	// Variables
	// utils.DeclaringASingleVariable()
	// utils.DeclaringAVariableWithAnInitialValue()
	// utils.MultipleVariableDeclaration()
	// utils.DeclareVariablesOfDifferentTypes()
	// utils.ShortHansDeclaration()
	// utils.VariablesAssignedValuesAndComputedDuringRuntime()

	//---------------------------------------------------
	//Types
	// utils.Bool()
	// utils.SignedIntegers()
	// utils.FloatingPointTypes()
	// utils.ComplessTypes()
	// utils.StringType()
	// utils.TypeConversion()

	//---------------------------------------------------
	//Constant
	// utils.DeclaringAConstant()
	// utils.DeclaringAGroupOfConstants()
	// utils.NemeriConstants()
	// 	utils.NumericExpressions()

	//---------------------------------------------------
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
	// area, _ := 10.8, 5.6
	// utils.RectProps3(area, 5.6)

	//---------------------------------------------------
	//Packages
	// utils.ImplementationExample()

	//---------------------------------------------------
	//Control Flow.. If else statement
	// utils.EvenOdd()
	// utils.EvenOdd1()
	// utils.LessOrGreater()
	// utils.EvenOdd2()
	// utils.IdiomaticEvenOdd()

	//---------------------------------------------------
	//Loops
	// utils.ForLoop()
	// utils.ForLoopBreak()
	// utils.ForLoopContinue()
	// utils.ForLoopNested()
	// utils.ForLoopLabels()
	// utils.ForLoopLabelsBreak()
	// utils.ForLoopLabelsOuter()
	// utils.MoreExamples()
	// utils.MoreExamplesSemicolonOmmitted()
	// utils.MultipleVariablesInForLoop()

	//---------------------------------------------------
	//Switch
	// utils.Switch()
	// utils.DefaultCase()
	// utils.MultipleExpressionInCase()
	// utils.ExpressionlessSwitch()
	// utils.Fallthrough()
	// utils.FallthroughHappensEvenWhenTheCaseEvaluatesToFalse()
	// utils.BreakingSwitch()
	// utils.BreakingTheOuterForLoop()

	//---------------------------------------------------
	//Array
	// utils.ArrayDeclaration1()
	// utils.ArrayDeclaration2()
	// utils.ArrayDeclaration3()
	// utils.ArrayDeclaration4()
	// utils.ArrayAreValueTypes()
	// utils.ImplementChangeLocalExample()
	// utils.LengthOfAnArray()
	// utils.IteratingArraysUsingRange1()
	// utils.IteratingArraysUsingRange2()

	//Multidimentional Arrays
	// utils.MultidimentionalArrays()

	//Slices
	// utils.Slice()
	// utils.SliceAnotherExampleAndScenario()
	// utils.ModifyingASlice()
	// utils.SliceAnotherExampleAndScenario2()
	// utils.LenghtAndCapacityOfASlice()
	// utils.CreatingASliceUsingMake()
	// utils.AppendingToASlice1()
	// utils.AppendingToASlice2()
	// utils.AppendingToASlice3()
	// utils.PassingASliceToAFunction()

	//Multidimentional Slices
	// utils.MultidimentionalSlices()
	// utils.MemoryOptimisation()

	//---------------------------------------------------
	//Variadic Function
	// `utils.Find1(89, 89, 90, 95)
	// utils.Find1(45, 56, 67, 45, 90, 109)
	// utils.Find1(78, 38, 56, 98)
	// utils.Find1(87)
	// utils.ImplementationPassingASliceToAVariadicFunction()
	// utils.Gotcha1()
	// utils.Gotcha2()

	//---------------------------------------------------
	//Maps
	// utils.CreateAMap()
	// utils.AddingItemsToAMap()
	// utils.InitializeAMapDuringDeclarationItself()
	// utils.RetrievingValueForAKeyFromAMap()
	// utils.IfElementIsNotPresentINAMap()
	// utils.CheckingIfAKeyExists()
	// utils.IterateOverAllElementsInAMap()
	// utils.DeletingItemsFromAMap()
	// utils.MapOfStructs()
	// utils.LengthOfTheMap()
	// utils.MapsAreReferenceTypes()

	// ---------------------------------------------------
	//Strings
	// utils.Strings()
	// utils.AccesingIndividualBytesOfAString()
	// utils.AccessingIndividualCharactersOfAString()
	// utils.Rune()
	// utils.AccessingIndividualRunesUsingForRangeLoop()
	// utils.CreatingAStringFromSliceOfBytes()
	// utils.DecimalEquivalentOfHexValues()
	// utils.CreatingAStringFromSliceOfRunes()
	// utils.StringLength()
	// utils.StringComparison()
	// utils.StringConcatenation()
	// utils.StringsAreImmutable()

	// ---------------------------------------------------
	// Pointers
	// utils.DeclaringPointers()
	// utils.ZeroValueOfApointer()
	// utils.CreatingPointersUsingTheNewFunction()
	// utils.DereferencingAPointer1()
	// utils.DereferencingAPointer2()
	// utils.PassingPointerToAFunction()
	// utils.ReturningPointerFromAFunctions()
	// utils.DoNotPassAPointerToAnArrayAsAnArgumentToAFunction1()
	// utils.DoNotPassAPointerToAnArrayAsAnArgumentToAFunction2()
	// utils.DoNotPassAPointerToAnArrayAsAnArgumentToAFunction3() // it is the best approach wheng  using array

	// ---------------------------------------------------
	//Structs
	// utils.DeclaringAStruct()
	// utils.CreatingAnonymousStructs()
	// utils.AccessingIndividualFieldsOfAStruct()
	// utils.ZeroValueOfAStruct()
	// utils.PointersToAStruct()
	// utils.AnonymousFields()
	// utils.NestedStructs()

	//exported structs
	// utils.ExportedStructsAndFields()
	// utils.StructsEquality()

	// ---------------------------------------------------
	//Methods
	// utils.CallingDisplaySalaryMethods()
	// utils.CallingDisplaySalaryFunction()
	// utils.CallingRectangleAndCicleTypeUsingMethod()
	// utils.CallingPointersRecieversVsValueReceivers()
	// utils.CallingMethodsOfAnonymousStructFields()
	// utils.CallingValueRecieversInMethodsVsValueArgumentsInFunctions()
	// utils.CallingPointerRecieversInMethodsVsPointerArgumentsInFunctions()
	// utils.MethodsWithNonStructReceivers()

	// ---------------------------------------------------
	//Interfaces1
	// utils.DeclaringAndImplementingAnInterface()
	// utils.PracticalUseOfAnInterface()
	// utils.PracticalUseOfAnInterface1()
	// utils.InterfaceInternalRepresentation()
	// utils.EmptyInterface()
	// utils.TypeAssertion1()
	// utils.TypeSwitch()
	// utils.TypeSwitch1()

	// ---------------------------------------------------
	//Interfaces2
	// utils.ImplementingInterfacesUsingPointerReceiversVsValueReceivers()
	// utils.ImplementingMultipleInterfaces()
	// utils.EmbeddingInterfaces()
	// utils.ZeroValueOfInterface()

	// ---------------------------------------------------
	// Concurrency
	// 1. Goroutines
	// utils.Goroutine1()
	// utils.StartingMultipleGoroutines()

	// 2. Channels
	// utils.DeclaringChannels()
	// utils.ChannelExampleProgram()
	// utils.AnotherExampleForChannels()
	// utils.UndirectionalChannels()
	// utils.ClosingChannelsAndForRangeLoopsOnChannels()
	// utils.AnotherExampleForChannels1()

	// 3. Buffered Channels
	// utils.BufferedChannelAndWorkerPoolsExample()
	// utils.WriteImplementation()
	// utils.Deadlock()
	// utils.ClosingBufferedChannels()
	// utils.RangeCloringBufferedChannels()
	// utils.LengthVsCapacity()
	// utils.WaitGroup()
	utils.WorkerPoolExample()
}
